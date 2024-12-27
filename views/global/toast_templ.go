// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package global

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ToastGlobalState() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_ToastGlobalState_7389`,
		Function: `function __templ_ToastGlobalState_7389(){document.addEventListener('alpine:init', () => {
		Alpine.store('toast', { 
			toastVisible: false,
			toastMessage: '',
			toastType: 'danger',
			
			displayToast(message, type) {
				this.toastVisible = true;
				this.toastMessage = message;
				this.toastType = type;
				setTimeout(() => {this.clearToast()}, 3000);
			},
 
			clearToast() {
				this.toastVisible = false;
				this.toastMessage = '';
			}    
		});
	});
 
	document.addEventListener('htmx:afterRequest', (event) => { 
		const contentType = event.detail.xhr.getResponseHeader("Content-Type");
		if (contentType !== 'application/json') {
			return;
		}
 
	    const responseData = event.detail.xhr.responseText;
		if (responseData === '') {
			return;
		}
 
		let toastType = 'success';
		const isResponseError = event.detail.xhr.status >= 400;
		if (isResponseError) {
			toastType = 'danger';
		}
 
		const parsedResponse = JSON.parse(responseData);
		if (parsedResponse.message === undefined || parsedResponse.message === '') {
			return;
		}
		const toastMessage = parsedResponse.message;   
	
		Alpine.store('toast').displayToast(toastMessage, toastType); 
	});
}`,
		Call:       templ.SafeScript(`__templ_ToastGlobalState_7389`),
		CallInline: templ.SafeScriptInline(`__templ_ToastGlobalState_7389`),
	}
}

func Toast() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = ToastGlobalState().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"modal\" :class=\"$store.toast.toastVisible ? &#39;is-active&#39; : &#39;&#39;\" x-data x-show=\"$store.toast.toastVisible\" x-transition x-cloak un-cloak><div class=\"modal-background\"></div><div class=\"modal-content\"><div class=\"box\" x-text=\"$store.toast.toastMessage\"></div></div><button class=\"modal-close is-large\" aria-label=\"close\" @click=\"$store.toast.clearToast()\"></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
