"use strict";(self.webpackChunkblogs=self.webpackChunkblogs||[]).push([[591],{3905:(e,r,t)=>{t.d(r,{Zo:()=>u,kt:()=>d});var n=t(7294);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function l(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function a(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?l(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):l(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function c(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},l=Object.keys(e);for(n=0;n<l.length;n++)t=l[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)t=l[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var i=n.createContext({}),p=function(e){var r=n.useContext(i),t=r;return e&&(t="function"==typeof e?e(r):a(a({},r),e)),t},u=function(e){var r=p(e.components);return n.createElement(i.Provider,{value:r},e.children)},s="mdxType",f={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},m=n.forwardRef((function(e,r){var t=e.components,o=e.mdxType,l=e.originalType,i=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),s=p(t),m=o,d=s["".concat(i,".").concat(m)]||s[m]||f[m]||l;return t?n.createElement(d,a(a({ref:r},u),{},{components:t})):n.createElement(d,a({ref:r},u))}));function d(e,r){var t=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var l=t.length,a=new Array(l);a[0]=m;var c={};for(var i in r)hasOwnProperty.call(r,i)&&(c[i]=r[i]);c.originalType=e,c[s]="string"==typeof e?e:o,a[1]=c;for(var p=2;p<l;p++)a[p]=t[p];return n.createElement.apply(null,a)}return n.createElement.apply(null,t)}m.displayName="MDXCreateElement"},4172:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>i,contentTitle:()=>a,default:()=>f,frontMatter:()=>l,metadata:()=>c,toc:()=>p});var n=t(7462),o=(t(7294),t(3905));const l={},a="Hello world",c={unversionedId:"hello-world/README",id:"hello-world/README",title:"Hello world",description:"main.go",source:"@site/../01-basic/01-hello-world/README.md",sourceDirName:"01-hello-world",slug:"/hello-world/",permalink:"/learn-go/hello-world/",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",next:{title:"Go basic",permalink:"/learn-go/"}},i={},p=[],u={toc:p},s="wrapper";function f(e){let{components:r,...t}=e;return(0,o.kt)(s,(0,n.Z)({},u,t,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"hello-world"},"Hello world"),(0,o.kt)("p",null,(0,o.kt)("inlineCode",{parentName:"p"},"main.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Print("Hello, World!")\n}\n')),(0,o.kt)("p",null,"\u0110\u1ec3 ch\u1ea1y code go ta d\xf9ng l\u1ec7nh ",(0,o.kt)("inlineCode",{parentName:"p"},"go run"),":"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"$ go run main.go\nHello, World!\n")))}f.isMDXComponent=!0}}]);