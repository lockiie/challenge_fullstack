(()=>{"use strict";var e,v={},g={};function r(e){var l=g[e];if(void 0!==l)return l.exports;var t=g[e]={exports:{}};return v[e](t,t.exports,r),t.exports}r.m=v,e=[],r.O=(l,t,i,u)=>{if(!t){var a=1/0;for(n=0;n<e.length;n++){for(var[t,i,u]=e[n],s=!0,f=0;f<t.length;f++)(!1&u||a>=u)&&Object.keys(r.O).every(b=>r.O[b](t[f]))?t.splice(f--,1):(s=!1,u<a&&(a=u));if(s){e.splice(n--,1);var d=i();void 0!==d&&(l=d)}}return l}u=u||0;for(var n=e.length;n>0&&e[n-1][2]>u;n--)e[n]=e[n-1];e[n]=[t,i,u]},r.n=e=>{var l=e&&e.__esModule?()=>e.default:()=>e;return r.d(l,{a:l}),l},r.d=(e,l)=>{for(var t in l)r.o(l,t)&&!r.o(e,t)&&Object.defineProperty(e,t,{enumerable:!0,get:l[t]})},r.f={},r.e=e=>Promise.all(Object.keys(r.f).reduce((l,t)=>(r.f[t](e,l),l),[])),r.u=e=>e+".1b7d6e9be9395ef5.js",r.miniCssF=e=>{},r.o=(e,l)=>Object.prototype.hasOwnProperty.call(e,l),(()=>{var e={},l="challenge-fullstack:";r.l=(t,i,u,n)=>{if(e[t])e[t].push(i);else{var a,s;if(void 0!==u)for(var f=document.getElementsByTagName("script"),d=0;d<f.length;d++){var o=f[d];if(o.getAttribute("src")==t||o.getAttribute("data-webpack")==l+u){a=o;break}}a||(s=!0,(a=document.createElement("script")).type="module",a.charset="utf-8",a.timeout=120,r.nc&&a.setAttribute("nonce",r.nc),a.setAttribute("data-webpack",l+u),a.src=r.tu(t)),e[t]=[i];var c=(m,b)=>{a.onerror=a.onload=null,clearTimeout(p);var h=e[t];if(delete e[t],a.parentNode&&a.parentNode.removeChild(a),h&&h.forEach(_=>_(b)),m)return m(b)},p=setTimeout(c.bind(null,void 0,{type:"timeout",target:a}),12e4);a.onerror=c.bind(null,a.onerror),a.onload=c.bind(null,a.onload),s&&document.head.appendChild(a)}}})(),r.r=e=>{"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},(()=>{var e;r.tu=l=>(void 0===e&&(e={createScriptURL:t=>t},"undefined"!=typeof trustedTypes&&trustedTypes.createPolicy&&(e=trustedTypes.createPolicy("angular#bundler",e))),e.createScriptURL(l))})(),r.p="",(()=>{var e={666:0};r.f.j=(i,u)=>{var n=r.o(e,i)?e[i]:void 0;if(0!==n)if(n)u.push(n[2]);else if(666!=i){var a=new Promise((o,c)=>n=e[i]=[o,c]);u.push(n[2]=a);var s=r.p+r.u(i),f=new Error;r.l(s,o=>{if(r.o(e,i)&&(0!==(n=e[i])&&(e[i]=void 0),n)){var c=o&&("load"===o.type?"missing":o.type),p=o&&o.target&&o.target.src;f.message="Loading chunk "+i+" failed.\n("+c+": "+p+")",f.name="ChunkLoadError",f.type=c,f.request=p,n[1](f)}},"chunk-"+i,i)}else e[i]=0},r.O.j=i=>0===e[i];var l=(i,u)=>{var f,d,[n,a,s]=u,o=0;if(n.some(p=>0!==e[p])){for(f in a)r.o(a,f)&&(r.m[f]=a[f]);if(s)var c=s(r)}for(i&&i(u);o<n.length;o++)r.o(e,d=n[o])&&e[d]&&e[d][0](),e[n[o]]=0;return r.O(c)},t=self.webpackChunkchallenge_fullstack=self.webpackChunkchallenge_fullstack||[];t.forEach(l.bind(null,0)),t.push=l.bind(null,t.push.bind(t))})()})();