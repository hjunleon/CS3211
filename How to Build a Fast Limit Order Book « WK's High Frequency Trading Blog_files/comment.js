var _____WB$wombat$assign$function_____ = function(name) {return (self._wb_wombat && self._wb_wombat.local_init && self._wb_wombat.local_init(name)) || self[name]; };
if (!self.__WB_pmw) { self.__WB_pmw = function(obj) { this.__WB_source = obj; return this; } }
{
  let window = _____WB$wombat$assign$function_____("window");
  let self = _____WB$wombat$assign$function_____("self");
  let document = _____WB$wombat$assign$function_____("document");
  let location = _____WB$wombat$assign$function_____("location");
  let top = _____WB$wombat$assign$function_____("top");
  let parent = _____WB$wombat$assign$function_____("parent");
  let frames = _____WB$wombat$assign$function_____("frames");
  let opener = _____WB$wombat$assign$function_____("opener");

(function(){function reply(authorId,commentId,commentBox){var author=MGJS.$(authorId).innerHTML;var insertStr='<a href="#'+commentId+'">@'+author.replace(/\t|\n|\r\n/g,"")+' </a> \n';appendReply(insertStr,commentBox);}
function quote(authorId,commentId,commentBodyId,commentBox){var author=MGJS.$(authorId).innerHTML;var comment=jQuery('#'+commentBodyId);comment.children('script').remove();comment.children('.pd-rating').remove();comment=comment.html();var insertStr='<blockquote cite="#'+commentBodyId+'">';insertStr+='\n<strong><a href="#'+commentId+'">'+author.replace(/\t|\n|\r\n/g,"")+'</a> :</strong>';insertStr+=comment.replace(/\t/g,"");insertStr+='</blockquote>\n';insertQuote(insertStr,commentBox);}
function appendReply(insertStr,commentBox){if(MGJS.$(commentBox)&&MGJS.$(commentBox).type=='textarea'){field=MGJS.$(commentBox);}else{alert("The comment box does not exist!");return false;}
if(field.value.indexOf(insertStr)>-1){alert("You've already appended this reply!");return false;}
if(field.value.replace(/\s|\t|\n/g,"")==''){field.value=insertStr;}else{field.value=field.value.replace(/[\n]*$/g,"")+'\n\n'+insertStr;}
field.focus();}
function insertQuote(insertStr,commentBox){if(MGJS.$(commentBox)&&MGJS.$(commentBox).type=='textarea'){field=MGJS.$(commentBox);}else{alert("The comment box does not exist!");return false;}
if(document.selection){field.focus();sel=document.selection.createRange();sel.text=insertStr;field.focus();}else if(field.selectionStart||field.selectionStart=='0'){var startPos=field.selectionStart;var endPos=field.selectionEnd;var cursorPos=startPos;field.value=field.value.substring(0,startPos)
+insertStr
+field.value.substring(endPos,field.value.length);cursorPos+=insertStr.length;field.focus();field.selectionStart=cursorPos;field.selectionEnd=cursorPos;}else{field.value+=insertStr;field.focus();}}
window['MGJS_CMT']={};window['MGJS_CMT']['reply']=reply;window['MGJS_CMT']['quote']=quote;})();

}
/*
     FILE ARCHIVED ON 15:47:49 Feb 19, 2011 AND RETRIEVED FROM THE
     INTERNET ARCHIVE ON 12:03:33 Feb 25, 2023.
     JAVASCRIPT APPENDED BY WAYBACK MACHINE, COPYRIGHT INTERNET ARCHIVE.

     ALL OTHER CONTENT MAY ALSO BE PROTECTED BY COPYRIGHT (17 U.S.C.
     SECTION 108(a)(3)).
*/
/*
playback timings (ms):
  captures_list: 112.518
  exclusion.robots: 0.081
  exclusion.robots.policy: 0.072
  cdx.remote: 0.062
  esindex: 0.01
  LoadShardBlock: 47.596 (3)
  PetaboxLoader3.datanode: 117.846 (5)
  load_resource: 211.095 (2)
  PetaboxLoader3.resolve: 122.435 (2)
*/