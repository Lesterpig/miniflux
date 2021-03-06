// Code generated by go generate; DO NOT EDIT.

package static // import "miniflux.app/ui/static"

var Javascripts = map[string]string{
	"app": `(function(){'use strict';class DomHelper{static isVisible(element){return element.offsetParent!==null;}
static openNewTab(url){let win=window.open("");win.opener=null;win.location=url;win.focus();}
static scrollPageTo(element){let windowScrollPosition=window.pageYOffset;let windowHeight=document.documentElement.clientHeight;let viewportPosition=windowScrollPosition+windowHeight;let itemBottomPosition=element.offsetTop+element.offsetHeight;if(viewportPosition-itemBottomPosition<0||viewportPosition-element.offsetTop>windowHeight){window.scrollTo(0,element.offsetTop-10);}}
static getVisibleElements(selector){let elements=document.querySelectorAll(selector);let result=[];for(let i=0;i<elements.length;i++){if(this.isVisible(elements[i])){result.push(elements[i]);}}
return result;}
static findParent(element,selector){for(;element&&element!==document;element=element.parentNode){if(element.classList.contains(selector)){return element;}}
return null;}
static hasPassiveEventListenerOption(){var passiveSupported=false;try{var options=Object.defineProperty({},"passive",{get:function(){passiveSupported=true;}});window.addEventListener("test",options,options);window.removeEventListener("test",options,options);}catch(err){passiveSupported=false;}
return passiveSupported;}}
class TouchHandler{constructor(){this.reset();}
reset(){this.touch={start:{x:-1,y:-1},move:{x:-1,y:-1},element:null};}
calculateDistance(){if(this.touch.start.x>=-1&&this.touch.move.x>=-1){let horizontalDistance=Math.abs(this.touch.move.x-this.touch.start.x);let verticalDistance=Math.abs(this.touch.move.y-this.touch.start.y);if(horizontalDistance>30&&verticalDistance<70){return this.touch.move.x-this.touch.start.x;}}
return 0;}
findElement(element){if(element.classList.contains("touch-item")){return element;}
return DomHelper.findParent(element,"touch-item");}
onTouchStart(event){if(event.touches===undefined||event.touches.length!==1){return;}
this.reset();this.touch.start.x=event.touches[0].clientX;this.touch.start.y=event.touches[0].clientY;this.touch.element=this.findElement(event.touches[0].target);}
onTouchMove(event){if(event.touches===undefined||event.touches.length!==1||this.element===null){return;}
this.touch.move.x=event.touches[0].clientX;this.touch.move.y=event.touches[0].clientY;let distance=this.calculateDistance();let absDistance=Math.abs(distance);if(absDistance>0){let opacity=1-(absDistance>75?0.9:absDistance/75*0.9);let tx=distance>75?75:(distance<-75?-75:distance);this.touch.element.style.opacity=opacity;this.touch.element.style.transform="translateX("+tx+"px)";event.preventDefault();}}
onTouchEnd(event){if(event.touches===undefined){return;}
if(this.touch.element!==null){let distance=Math.abs(this.calculateDistance());if(distance>75){toggleEntryStatus(this.touch.element);}
this.touch.element.style.opacity=1;this.touch.element.style.transform="none";}
this.reset();}
listen(){let elements=document.querySelectorAll(".touch-item");let hasPassiveOption=DomHelper.hasPassiveEventListenerOption();elements.forEach((element)=>{element.addEventListener("touchstart",(e)=>this.onTouchStart(e),hasPassiveOption?{passive:true}:false);element.addEventListener("touchmove",(e)=>this.onTouchMove(e),hasPassiveOption?{passive:false}:false);element.addEventListener("touchend",(e)=>this.onTouchEnd(e),hasPassiveOption?{passive:true}:false);element.addEventListener("touchcancel",()=>this.reset(),hasPassiveOption?{passive:true}:false);});let entryContentElement=document.querySelector(".entry-content");if(entryContentElement){let doubleTapTimers={previous:null,next:null};const detectDoubleTap=(doubleTapTimer,event)=>{const timer=doubleTapTimers[doubleTapTimer];if(timer===null){doubleTapTimers[doubleTapTimer]=setTimeout(()=>{doubleTapTimers[doubleTapTimer]=null;},200);}else{event.preventDefault();goToPage(doubleTapTimer);}};entryContentElement.addEventListener("touchend",(e)=>{if(e.changedTouches[0].clientX>=(entryContentElement.offsetWidth/2)){detectDoubleTap("next",e);}else{detectDoubleTap("previous",e);}},hasPassiveOption?{passive:false}:false);entryContentElement.addEventListener("touchmove",(e)=>{Object.keys(doubleTapTimers).forEach(timer=>doubleTapTimers[timer]=null);});}}}
class KeyboardHandler{constructor(){this.queue=[];this.shortcuts={};}
on(combination,callback){this.shortcuts[combination]=callback;}
listen(){document.onkeydown=(event)=>{if(this.isEventIgnored(event)||this.isModifierKeyDown(event)){return;}
let key=this.getKey(event);this.queue.push(key);for(let combination in this.shortcuts){let keys=combination.split(" ");if(keys.every((value,index)=>value===this.queue[index])){this.queue=[];this.shortcuts[combination](event);return;}
if(keys.length===1&&key===keys[0]){this.queue=[];this.shortcuts[combination](event);return;}}
if(this.queue.length>=2){this.queue=[];}};}
isEventIgnored(event){return event.target.tagName==="INPUT"||event.target.tagName==="TEXTAREA";}
isModifierKeyDown(event){return event.getModifierState("Control")||event.getModifierState("Alt")||event.getModifierState("Meta");}
getKey(event){const mapping={'Esc':'Escape','Up':'ArrowUp','Down':'ArrowDown','Left':'ArrowLeft','Right':'ArrowRight'};for(let key in mapping){if(mapping.hasOwnProperty(key)&&key===event.key){return mapping[key];}}
return event.key;}}
class RequestBuilder{constructor(url){this.callback=null;this.url=url;this.options={method:"POST",cache:"no-cache",credentials:"include",body:null,headers:new Headers({"Content-Type":"application/json","X-Csrf-Token":this.getCsrfToken()})};}
withBody(body){this.options.body=JSON.stringify(body);return this;}
withCallback(callback){this.callback=callback;return this;}
getCsrfToken(){let element=document.querySelector("meta[name=X-CSRF-Token]");if(element!==null){return element.getAttribute("value");}
return "";}
execute(){fetch(new Request(this.url,this.options)).then((response)=>{if(this.callback){this.callback(response);}});}}
class ModalHandler{static exists(){return document.getElementById("modal-container")!==null;}
static open(fragment){if(ModalHandler.exists()){return;}
let container=document.createElement("div");container.id="modal-container";container.appendChild(document.importNode(fragment,true));document.body.appendChild(container);let closeButton=document.querySelector("a.btn-close-modal");if(closeButton!==null){closeButton.onclick=(event)=>{event.preventDefault();ModalHandler.close();};}}
static close(){let container=document.getElementById("modal-container");if(container!==null){container.parentNode.removeChild(container);}}}
function onClick(selector,callback,noPreventDefault){let elements=document.querySelectorAll(selector);elements.forEach((element)=>{element.onclick=(event)=>{if(!noPreventDefault){event.preventDefault();}
callback(event);};});}
function toggleMainMenu(){let menu=document.querySelector(".header nav ul");if(DomHelper.isVisible(menu)){menu.style.display="none";}else{menu.style.display="block";}
let searchElement=document.querySelector(".header .search");if(DomHelper.isVisible(searchElement)){searchElement.style.display="none";}else{searchElement.style.display="block";}}
function onClickMainMenuListItem(event){let element=event.target;if(element.tagName==="A"){window.location.href=element.getAttribute("href");}else{window.location.href=element.querySelector("a").getAttribute("href");}}
function handleSubmitButtons(){let elements=document.querySelectorAll("form");elements.forEach((element)=>{element.onsubmit=()=>{let button=element.querySelector("button");if(button){button.innerHTML=button.dataset.labelLoading;button.disabled=true;}};});}
function setFocusToSearchInput(event){event.preventDefault();event.stopPropagation();let toggleSwitchElement=document.querySelector(".search-toggle-switch");if(toggleSwitchElement){toggleSwitchElement.style.display="none";}
let searchFormElement=document.querySelector(".search-form");if(searchFormElement){searchFormElement.style.display="block";}
let searchInputElement=document.getElementById("search-input");if(searchInputElement){searchInputElement.focus();searchInputElement.value="";}}
function showKeyboardShortcuts(){let template=document.getElementById("keyboard-shortcuts");if(template!==null){ModalHandler.open(template.content);}}
function markPageAsRead(){let items=DomHelper.getVisibleElements(".items .item");let entryIDs=[];items.forEach((element)=>{element.classList.add("item-status-read");entryIDs.push(parseInt(element.dataset.id,10));});if(entryIDs.length>0){updateEntriesStatus(entryIDs,"read",()=>{let element=document.querySelector("a[data-action=markPageAsRead]");let showOnlyUnread=false;if(element){showOnlyUnread=element.dataset.showOnlyUnread||false;}
if(showOnlyUnread){window.location.reload();}else{goToPage("next",true);}});}}
function handleEntryStatus(element){let currentEntry=findEntry(element);if(currentEntry){toggleEntryStatus(currentEntry);if(isListView()&&currentEntry.classList.contains('current-item')){goToNextListItem();}}}
function toggleEntryStatus(element){let entryID=parseInt(element.dataset.id,10);let link=element.querySelector("a[data-toggle-status]");let currentStatus=link.dataset.value;let newStatus=currentStatus==="read"?"unread":"read";updateEntriesStatus([entryID],newStatus);if(currentStatus==="read"){link.innerHTML=link.dataset.labelRead;link.dataset.value="unread";}else{link.innerHTML=link.dataset.labelUnread;link.dataset.value="read";}
if(element.classList.contains("item-status-"+currentStatus)){element.classList.remove("item-status-"+currentStatus);element.classList.add("item-status-"+newStatus);}}
function markEntryAsRead(element){if(element.classList.contains("item-status-unread")){element.classList.remove("item-status-unread");element.classList.add("item-status-read");let entryID=parseInt(element.dataset.id,10);updateEntriesStatus([entryID],"read");}}
function updateEntriesStatus(entryIDs,status,callback){let url=document.body.dataset.entriesStatusUrl;let request=new RequestBuilder(url);request.withBody({entry_ids:entryIDs,status:status});request.withCallback(callback);request.execute();if(status==="read"){decrementUnreadCounter(1);}else{incrementUnreadCounter(1);}}
function handleSaveEntry(element){let currentEntry=findEntry(element);if(currentEntry){saveEntry(currentEntry.querySelector("a[data-save-entry]"));}}
function saveEntry(element){if(!element){return;}
if(element.dataset.completed){return;}
element.innerHTML=element.dataset.labelLoading;let request=new RequestBuilder(element.dataset.saveUrl);request.withCallback(()=>{element.innerHTML=element.dataset.labelDone;element.dataset.completed=true;});request.execute();}
function handleBookmark(element){let currentEntry=findEntry(element);if(currentEntry){toggleBookmark(currentEntry);}}
function toggleBookmark(parentElement){let element=parentElement.querySelector("a[data-toggle-bookmark]");if(!element){return;}
element.innerHTML=element.dataset.labelLoading;let request=new RequestBuilder(element.dataset.bookmarkUrl);request.withCallback(()=>{if(element.dataset.value==="star"){element.innerHTML=element.dataset.labelStar;element.dataset.value="unstar";}else{element.innerHTML=element.dataset.labelUnstar;element.dataset.value="star";}});request.execute();}
function handleFetchOriginalContent(){if(isListView()){return;}
let element=document.querySelector("a[data-fetch-content-entry]");if(!element){return;}
if(element.dataset.completed){return;}
element.innerHTML=element.dataset.labelLoading;let request=new RequestBuilder(element.dataset.fetchContentUrl);request.withCallback((response)=>{element.innerHTML=element.dataset.labelDone;element.dataset.completed=true;response.json().then((data)=>{if(data.hasOwnProperty("content")){document.querySelector(".entry-content").innerHTML=data.content;}});});request.execute();}
function openOriginalLink(){let entryLink=document.querySelector(".entry h1 a");if(entryLink!==null){DomHelper.openNewTab(entryLink.getAttribute("href"));return;}
let currentItemOriginalLink=document.querySelector(".current-item a[data-original-link]");if(currentItemOriginalLink!==null){DomHelper.openNewTab(currentItemOriginalLink.getAttribute("href"));let currentItem=document.querySelector(".current-item");goToNextListItem();markEntryAsRead(currentItem);}}
function openSelectedItem(){let currentItemLink=document.querySelector(".current-item .item-title a");if(currentItemLink!==null){window.location.href=currentItemLink.getAttribute("href");}}
function unsubscribeFromFeed(){let unsubscribeLinks=document.querySelectorAll("[data-action=remove-feed]");if(unsubscribeLinks.length===1){let unsubscribeLink=unsubscribeLinks[0];let request=new RequestBuilder(unsubscribeLink.dataset.url);request.withCallback(()=>{if(unsubscribeLink.dataset.redirectUrl){window.location.href=unsubscribeLink.dataset.redirectUrl;}else{window.location.reload();}});request.execute();}}
function goToPage(page,fallbackSelf){let element=document.querySelector("a[data-page="+page+"]");if(element){document.location.href=element.href;}else if(fallbackSelf){window.location.reload();}}
function goToPrevious(){if(isListView()){goToPreviousListItem();}else{goToPage("previous");}}
function goToNext(){if(isListView()){goToNextListItem();}else{goToPage("next");}}
function goToFeedOrFeeds(){if(isEntry()){let feedAnchor=document.querySelector("span.entry-website a");if(feedAnchor!==null){window.location.href=feedAnchor.href;}}else{goToPage('feeds');}}
function goToPreviousListItem(){let items=DomHelper.getVisibleElements(".items .item");if(items.length===0){return;}
if(document.querySelector(".current-item")===null){items[0].classList.add("current-item");items[0].querySelector('.item-header a').focus();return;}
for(let i=0;i<items.length;i++){if(items[i].classList.contains("current-item")){items[i].classList.remove("current-item");if(i-1>=0){items[i-1].classList.add("current-item");DomHelper.scrollPageTo(items[i-1]);items[i-1].querySelector('.item-header a').focus();}
break;}}}
function goToNextListItem(){let currentItem=document.querySelector(".current-item");let items=DomHelper.getVisibleElements(".items .item");if(items.length===0){return;}
if(currentItem===null){items[0].classList.add("current-item");items[0].querySelector('.item-header a').focus();return;}
for(let i=0;i<items.length;i++){if(items[i].classList.contains("current-item")){items[i].classList.remove("current-item");if(i+1<items.length){items[i+1].classList.add("current-item");DomHelper.scrollPageTo(items[i+1]);items[i+1].querySelector('.item-header a').focus();}
break;}}}
function decrementUnreadCounter(n){updateUnreadCounterValue((current)=>{return current-n;});}
function incrementUnreadCounter(n){updateUnreadCounterValue((current)=>{return current+n;});}
function updateUnreadCounterValue(callback){let counterElements=document.querySelectorAll("span.unread-counter");counterElements.forEach((element)=>{let oldValue=parseInt(element.textContent,10);element.innerHTML=callback(oldValue);});if(window.location.href.endsWith('/unread')){let oldValue=parseInt(document.title.split('(')[1],10);let newValue=callback(oldValue);document.title=document.title.replace(/(.*?)\(\d+\)(.*?)/,function(match,prefix,suffix,offset,string){return prefix+'('+newValue+')'+suffix;});}}
function isEntry(){return document.querySelector("section.entry")!==null;}
function isListView(){return document.querySelector(".items")!==null;}
function findEntry(element){if(isListView()){if(element){return DomHelper.findParent(element,"item");}else{return document.querySelector(".current-item");}}else{return document.querySelector(".entry");}}
function handleConfirmationMessage(linkElement,callback){linkElement.style.display="none";let containerElement=linkElement.parentNode;let questionElement=document.createElement("span");let yesElement=document.createElement("a");yesElement.href="#";yesElement.appendChild(document.createTextNode(linkElement.dataset.labelYes));yesElement.onclick=(event)=>{event.preventDefault();let loadingElement=document.createElement("span");loadingElement.className="loading";loadingElement.appendChild(document.createTextNode(linkElement.dataset.labelLoading));questionElement.remove();containerElement.appendChild(loadingElement);callback(linkElement.dataset.url,linkElement.dataset.redirectUrl);};let noElement=document.createElement("a");noElement.href="#";noElement.appendChild(document.createTextNode(linkElement.dataset.labelNo));noElement.onclick=(event)=>{event.preventDefault();linkElement.style.display="inline";questionElement.remove();};questionElement.className="confirm";questionElement.appendChild(document.createTextNode(linkElement.dataset.labelQuestion+" "));questionElement.appendChild(yesElement);questionElement.appendChild(document.createTextNode(", "));questionElement.appendChild(noElement);containerElement.appendChild(questionElement);}
document.addEventListener("DOMContentLoaded",function(){handleSubmitButtons();if(!document.querySelector("body[data-disable-keyboard-shortcuts=true]")){let keyboardHandler=new KeyboardHandler();keyboardHandler.on("g u",()=>goToPage("unread"));keyboardHandler.on("g b",()=>goToPage("starred"));keyboardHandler.on("g h",()=>goToPage("history"));keyboardHandler.on("g f",()=>goToFeedOrFeeds());keyboardHandler.on("g c",()=>goToPage("categories"));keyboardHandler.on("g s",()=>goToPage("settings"));keyboardHandler.on("ArrowLeft",()=>goToPrevious());keyboardHandler.on("ArrowRight",()=>goToNext());keyboardHandler.on("k",()=>goToPrevious());keyboardHandler.on("p",()=>goToPrevious());keyboardHandler.on("j",()=>goToNext());keyboardHandler.on("n",()=>goToNext());keyboardHandler.on("h",()=>goToPage("previous"));keyboardHandler.on("l",()=>goToPage("next"));keyboardHandler.on("o",()=>openSelectedItem());keyboardHandler.on("v",()=>openOriginalLink());keyboardHandler.on("m",()=>handleEntryStatus());keyboardHandler.on("A",()=>markPageAsRead());keyboardHandler.on("s",()=>handleSaveEntry());keyboardHandler.on("d",()=>handleFetchOriginalContent());keyboardHandler.on("f",()=>handleBookmark());keyboardHandler.on("?",()=>showKeyboardShortcuts());keyboardHandler.on("#",()=>unsubscribeFromFeed());keyboardHandler.on("/",(e)=>setFocusToSearchInput(e));keyboardHandler.on("Escape",()=>ModalHandler.close());keyboardHandler.listen();}
let touchHandler=new TouchHandler();touchHandler.listen();onClick("a[data-save-entry]",(event)=>handleSaveEntry(event.target));onClick("a[data-toggle-bookmark]",(event)=>handleBookmark(event.target));onClick("a[data-fetch-content-entry]",()=>handleFetchOriginalContent());onClick("a[data-action=search]",(event)=>setFocusToSearchInput(event));onClick("a[data-action=markPageAsRead]",()=>handleConfirmationMessage(event.target,()=>markPageAsRead()));onClick("a[data-toggle-status]",(event)=>handleEntryStatus(event.target));onClick("a[data-confirm]",(event)=>handleConfirmationMessage(event.target,(url,redirectURL)=>{let request=new RequestBuilder(url);request.withCallback(()=>{if(redirectURL){window.location.href=redirectURL;}else{window.location.reload();}});request.execute();}));if(document.documentElement.clientWidth<600){onClick(".logo",()=>toggleMainMenu());onClick(".header nav li",(event)=>onClickMainMenuListItem(event));}
if("serviceWorker"in navigator){let scriptElement=document.getElementById("service-worker-script");if(scriptElement){navigator.serviceWorker.register(scriptElement.src);}}});})();`,
	"sw": `'use strict';self.addEventListener("fetch",(event)=>{if(event.request.url.includes("/feed/icon/")){event.respondWith(caches.open("feed_icons").then((cache)=>{return cache.match(event.request).then((response)=>{return response||fetch(event.request).then((response)=>{cache.put(event.request,response.clone());return response;});});}));}});`,
}

var JavascriptsChecksums = map[string]string{
	"app": "3e73bb4f1be3c679e59dcf91560efdf8646d7d549e682e5ad83f3bb1a6eeeff7",
	"sw":  "55fffa223919cc18572788fb9c62fccf92166c0eb5d3a1d6f91c31f24d020be9",
}
