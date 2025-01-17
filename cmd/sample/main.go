package main

import (
	"io/fs"
	"os"
	resources "simple-server"
	"simple-server/internal"
	"simple-server/projects/sample/handlers"
	"simple-server/projects/sample/jobs"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	/* 환경 설정 */
	internal.LoadEnv()
	_ = os.Setenv("APP_NAME", "Sample")
	_ = os.Setenv("APP_DATABASE_URL", "file:./projects/sample/pb_data/data.db")
	_ = os.Setenv("LOG_DATABASE_URL", "file:./projects/sample/pb_data/auxiliary.db")
	/* 환경 설정 */

	/* 로깅 초기화 */
	internal.LoggerWithDatabase()
	/* 로깅 초기화 */

	/* 파이어베이스 초기화 */
	internal.FirebaseInit()
	/* 파이어베이스 초기화 */

	e := echo.New()

	/* 미들 웨어 */
	sharedStaticFS, _ := fs.Sub(resources.EmbeddedFiles, "shared/static")
	projectStaticFS, _ := fs.Sub(resources.EmbeddedFiles, "projects/sample/static")
	e.StaticFS("/shared/static", sharedStaticFS) // 공통 정적 파일
	e.StaticFS("/static", projectStaticFS)       // 프로젝트 정적 파일
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("sample"))
	e.GET("/metrics", echoprometheus.NewHandler())

	// 공개 그룹
	public := e.Group("")

	// 인증 그룹
	private := e.Group("")

	private.Use(middleware.KeyAuthWithConfig(internal.FirebaseAuth()))
	/* 미들 웨어 */

	/* 라우터  */
	public.GET("/", handlers.IndexPageHandler)
	public.GET("/login", handlers.LoginPageHanlder)
	public.GET("/squash", func(c echo.Context) error { // 스쿼시 잡 실행
		jobs.SquashExecute()
		return c.String(200, "Squash 실행")
	})

	private.GET("/authors", handlers.GetAuthors)     // 저자 리스트 조회
	private.GET("/author", handlers.GetAuthor)       // 저자 조회
	private.POST("/author", handlers.CreateAuthor)   // 저자 등록
	private.PUT("/author", handlers.UpdateAuthor)    // 저자 수정
	private.DELETE("/author", handlers.DeleteAuthor) // 저자 삭제

	private.GET("/reset-form", handlers.ResetForm) // 저자 등록폼 리셋
	/* 라우터  */

	/* 크론 잡 */
	// c := cron.New()

	// jobs.SquashJob(c)

	// go func() {
	// 	c.Start()
	// }()
	/* 크론 잡 */

	e.Logger.Fatal(e.Start(":8000"))
}
