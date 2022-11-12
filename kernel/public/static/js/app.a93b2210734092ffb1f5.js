webpackJsonp([1],{LUVb:function(t,e){},NHnr:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=n("7+uW"),a={data:function(){return{dialogFormVisible:!1,tableData:null,applist:null,form:{SoftName:"",SoftVersion:""},formInline:{SoftName:"",SoftPath:"",SoftVersion:""},activeName:"first"}},beforeCreate:function(){document.querySelector("body").setAttribute("style","background-color:#525252;")},created:function(){var t=this;this.$axios.post("http://127.0.0.1:8899/get_lists").then(function(e){t.tableData=e.data.data.list,t.applist=e.data.data.list}).catch(function(t){return console.log(t)})},methods:{deleteRow:function(t,e){var n=this,o=new URLSearchParams;o.append("name",e[t].SoftName),o.append("version",e[t].SoftVersion),this.$axios.post("http://127.0.0.1:8899/delete",o).then(function(o){e.splice(t,1),n.$notify({title:"移除成功",type:"success"})}).catch(function(t){return console.log(t)})},open1:function(){this.$notify({title:"已终止",message:"您已成功终止程序运行",type:"success"})},change:function(){var t=this;if(this.form.SoftName){var e=new URLSearchParams,n=this.form.SoftName.split("-",2);e.append("name",n[0]),e.append("version",n[1]),this.$axios.post("http://127.0.0.1:8899/switch",e).then(function(e){t.$notify({title:"切换成功",message:"已切换"+n[0]+"-"+n[1],type:"success"})}).catch(function(t){return console.log(t)})}},DoAdd:function(){var t=this;if(this.formInline.SoftName&&this.formInline.SoftVersion&&this.formInline.SoftPath){var e=new URLSearchParams;e.append("name",this.formInline.SoftName),e.append("version",this.formInline.SoftVersion),e.append("path",this.formInline.SoftPath),this.$axios.post("http://127.0.0.1:8899/create",e).then(function(e){t.$notify({title:"新增成功",message:"已新增"+t.formInline.SoftName+"-"+t.formInline.SoftVersion+"路径："+t.formInline.SoftPath,type:"success"}),t.formInline=[],t.dialogFormVisible=!1}).catch(function(t){return console.log(t)})}},StopApp:function(){this.open1()},handleClick:function(t,e){console.log(t,e)}}},l={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"app"}},[n("div",[n("el-button",{attrs:{type:"primary",plain:""}},[t._v("Switch Env V1.0")])],1),t._v(" "),n("el-tabs",{on:{"tab-click":t.handleClick},model:{value:t.activeName,callback:function(e){t.activeName=e},expression:"activeName"}},[n("el-tab-pane",{attrs:{label:"软件配置",name:"first"}},[n("div",{staticStyle:{"background-color":"#fff","box-shadow":"0px 4px 10px darkgrey",padding:"5px","text-align":"center"}},[n("div",{staticStyle:{height:"1rem"}}),t._v(" "),n("div",{staticStyle:{height:"1rem"}}),t._v(" "),n("el-form",{ref:"form",attrs:{model:t.form}},[n("el-form-item",{attrs:{label:"切换版本"}},[n("el-select",{attrs:{placeholder:"请选择软件版本"},model:{value:t.form.SoftName,callback:function(e){t.$set(t.form,"SoftName",e)},expression:"form.SoftName"}},t._l(t.applist,function(t,e){return n("el-option",{key:e,attrs:{label:t.key,value:t.SoftName+"-"+t.SoftVersion}})}),1)],1),t._v(" "),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:t.change}},[t._v("切换")])],1)],1)],1)]),t._v(" "),n("el-tab-pane",{attrs:{label:"版本管理",name:"second"}},[n("el-button",{staticStyle:{float:"left"},attrs:{type:"primary",size:"small"},on:{click:function(e){t.dialogFormVisible=!0}}},[t._v("新增安装包")]),t._v(" "),n("el-dialog",{attrs:{title:"新增软件",visible:t.dialogFormVisible},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[n("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:t.formInline}},[n("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"请务必保持小写：php",placement:"top-start"}},[n("el-form-item",{attrs:{label:"软件名称"}},[n("el-input",{attrs:{placeholder:"例如：php"},model:{value:t.formInline.SoftName,callback:function(e){t.$set(t.formInline,"SoftName",e)},expression:"formInline.SoftName"}})],1)],1),t._v(" "),n("el-form-item",{attrs:{label:"版本型号"}},[n("el-input",{attrs:{placeholder:"例如：74"},model:{value:t.formInline.SoftVersion,callback:function(e){t.$set(t.formInline,"SoftVersion",e)},expression:"formInline.SoftVersion"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"所在路径"}},[n("el-input",{attrs:{placeholder:"例如：D:/Daixs/php73"},model:{value:t.formInline.SoftPath,callback:function(e){t.$set(t.formInline,"SoftPath",e)},expression:"formInline.SoftPath"}})],1),t._v(" "),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:t.DoAdd}},[t._v("保存")])],1)],1)],1),t._v(" "),n("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData}},[n("el-table-column",{attrs:{label:"软件名称"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",{staticStyle:{"margin-left":"10px"}},[t._v(t._s(e.row.SoftName)+"-"+t._s(e.row.SoftVersion))])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:"软件路径"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",{staticStyle:{"margin-left":"10px"}},[t._v(t._s(e.row.SoftPath))])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{size:"mini",type:"danger"},nativeOn:{click:function(n){return n.preventDefault(),t.deleteRow(e.$index,t.tableData)}}},[t._v("移除")])]}}])})],1)],1),t._v(" "),n("el-tab-pane",{attrs:{label:"程序设置",name:"third"}},[n("el-button",{attrs:{type:"danger"},on:{click:t.StopApp}},[t._v("终止程序")])],1)],1)],1)},staticRenderFns:[]};var i=n("VU/8")(a,l,!1,function(t){n("LUVb")},null,null).exports,r=n("zL8q"),s=n.n(r),f=(n("tvR6"),n("mtWM"));o.default.config.productionTip=!1,o.default.prototype.$axios=f.a,o.default.use(s.a),new o.default({el:"#app",render:function(t){return t(i)}})},tvR6:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.a93b2210734092ffb1f5.js.map