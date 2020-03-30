(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["main"],{

/***/ "../../node_modules/tslib/tslib.es6.js":
/*!*******************************************************************!*\
  !*** /Users/wonder1/go/src/pi/js/node_modules/tslib/tslib.es6.js ***!
  \*******************************************************************/
/*! exports provided: __extends, __assign, __rest, __decorate, __param, __metadata, __awaiter, __generator, __exportStar, __values, __read, __spread, __spreadArrays, __await, __asyncGenerator, __asyncDelegator, __asyncValues, __makeTemplateObject, __importStar, __importDefault */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__extends", function() { return __extends; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__assign", function() { return __assign; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__rest", function() { return __rest; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__decorate", function() { return __decorate; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__param", function() { return __param; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__metadata", function() { return __metadata; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__awaiter", function() { return __awaiter; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__generator", function() { return __generator; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__exportStar", function() { return __exportStar; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__values", function() { return __values; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__read", function() { return __read; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__spread", function() { return __spread; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__spreadArrays", function() { return __spreadArrays; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__await", function() { return __await; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__asyncGenerator", function() { return __asyncGenerator; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__asyncDelegator", function() { return __asyncDelegator; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__asyncValues", function() { return __asyncValues; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__makeTemplateObject", function() { return __makeTemplateObject; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__importStar", function() { return __importStar; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "__importDefault", function() { return __importDefault; });
/*! *****************************************************************************
Copyright (c) Microsoft Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at http://www.apache.org/licenses/LICENSE-2.0

THIS CODE IS PROVIDED ON AN *AS IS* BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, EITHER EXPRESS OR IMPLIED, INCLUDING WITHOUT LIMITATION ANY IMPLIED
WARRANTIES OR CONDITIONS OF TITLE, FITNESS FOR A PARTICULAR PURPOSE,
MERCHANTABLITY OR NON-INFRINGEMENT.

See the Apache Version 2.0 License for specific language governing permissions
and limitations under the License.
***************************************************************************** */
/* global Reflect, Promise */

var extendStatics = function(d, b) {
    extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return extendStatics(d, b);
};

function __extends(d, b) {
    extendStatics(d, b);
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
}

var __assign = function() {
    __assign = Object.assign || function __assign(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p)) t[p] = s[p];
        }
        return t;
    }
    return __assign.apply(this, arguments);
}

function __rest(s, e) {
    var t = {};
    for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p) && e.indexOf(p) < 0)
        t[p] = s[p];
    if (s != null && typeof Object.getOwnPropertySymbols === "function")
        for (var i = 0, p = Object.getOwnPropertySymbols(s); i < p.length; i++) {
            if (e.indexOf(p[i]) < 0 && Object.prototype.propertyIsEnumerable.call(s, p[i]))
                t[p[i]] = s[p[i]];
        }
    return t;
}

function __decorate(decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
}

function __param(paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
}

function __metadata(metadataKey, metadataValue) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(metadataKey, metadataValue);
}

function __awaiter(thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
}

function __generator(thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
}

function __exportStar(m, exports) {
    for (var p in m) if (!exports.hasOwnProperty(p)) exports[p] = m[p];
}

function __values(o) {
    var m = typeof Symbol === "function" && o[Symbol.iterator], i = 0;
    if (m) return m.call(o);
    return {
        next: function () {
            if (o && i >= o.length) o = void 0;
            return { value: o && o[i++], done: !o };
        }
    };
}

function __read(o, n) {
    var m = typeof Symbol === "function" && o[Symbol.iterator];
    if (!m) return o;
    var i = m.call(o), r, ar = [], e;
    try {
        while ((n === void 0 || n-- > 0) && !(r = i.next()).done) ar.push(r.value);
    }
    catch (error) { e = { error: error }; }
    finally {
        try {
            if (r && !r.done && (m = i["return"])) m.call(i);
        }
        finally { if (e) throw e.error; }
    }
    return ar;
}

function __spread() {
    for (var ar = [], i = 0; i < arguments.length; i++)
        ar = ar.concat(__read(arguments[i]));
    return ar;
}

function __spreadArrays() {
    for (var s = 0, i = 0, il = arguments.length; i < il; i++) s += arguments[i].length;
    for (var r = Array(s), k = 0, i = 0; i < il; i++)
        for (var a = arguments[i], j = 0, jl = a.length; j < jl; j++, k++)
            r[k] = a[j];
    return r;
};

function __await(v) {
    return this instanceof __await ? (this.v = v, this) : new __await(v);
}

function __asyncGenerator(thisArg, _arguments, generator) {
    if (!Symbol.asyncIterator) throw new TypeError("Symbol.asyncIterator is not defined.");
    var g = generator.apply(thisArg, _arguments || []), i, q = [];
    return i = {}, verb("next"), verb("throw"), verb("return"), i[Symbol.asyncIterator] = function () { return this; }, i;
    function verb(n) { if (g[n]) i[n] = function (v) { return new Promise(function (a, b) { q.push([n, v, a, b]) > 1 || resume(n, v); }); }; }
    function resume(n, v) { try { step(g[n](v)); } catch (e) { settle(q[0][3], e); } }
    function step(r) { r.value instanceof __await ? Promise.resolve(r.value.v).then(fulfill, reject) : settle(q[0][2], r); }
    function fulfill(value) { resume("next", value); }
    function reject(value) { resume("throw", value); }
    function settle(f, v) { if (f(v), q.shift(), q.length) resume(q[0][0], q[0][1]); }
}

function __asyncDelegator(o) {
    var i, p;
    return i = {}, verb("next"), verb("throw", function (e) { throw e; }), verb("return"), i[Symbol.iterator] = function () { return this; }, i;
    function verb(n, f) { i[n] = o[n] ? function (v) { return (p = !p) ? { value: __await(o[n](v)), done: n === "return" } : f ? f(v) : v; } : f; }
}

function __asyncValues(o) {
    if (!Symbol.asyncIterator) throw new TypeError("Symbol.asyncIterator is not defined.");
    var m = o[Symbol.asyncIterator], i;
    return m ? m.call(o) : (o = typeof __values === "function" ? __values(o) : o[Symbol.iterator](), i = {}, verb("next"), verb("throw"), verb("return"), i[Symbol.asyncIterator] = function () { return this; }, i);
    function verb(n) { i[n] = o[n] && function (v) { return new Promise(function (resolve, reject) { v = o[n](v), settle(resolve, reject, v.done, v.value); }); }; }
    function settle(resolve, reject, d, v) { Promise.resolve(v).then(function(v) { resolve({ value: v, done: d }); }, reject); }
}

function __makeTemplateObject(cooked, raw) {
    if (Object.defineProperty) { Object.defineProperty(cooked, "raw", { value: raw }); } else { cooked.raw = raw; }
    return cooked;
};

function __importStar(mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result.default = mod;
    return result;
}

function __importDefault(mod) {
    return (mod && mod.__esModule) ? mod : { default: mod };
}


/***/ }),

/***/ "./$$_lazy_route_resource lazy recursive":
/*!******************************************************!*\
  !*** ./$$_lazy_route_resource lazy namespace object ***!
  \******************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

function webpackEmptyAsyncContext(req) {
	// Here Promise.resolve().then() is used instead of new Promise() to prevent
	// uncaught exception popping up in devtools
	return Promise.resolve().then(function() {
		var e = new Error("Cannot find module '" + req + "'");
		e.code = 'MODULE_NOT_FOUND';
		throw e;
	});
}
webpackEmptyAsyncContext.keys = function() { return []; };
webpackEmptyAsyncContext.resolve = webpackEmptyAsyncContext;
module.exports = webpackEmptyAsyncContext;
webpackEmptyAsyncContext.id = "./$$_lazy_route_resource lazy recursive";

/***/ }),

/***/ "./src/app/app.component.scss":
/*!************************************!*\
  !*** ./src/app/app.component.scss ***!
  \************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ("/* sidebar */\n.sidebar {\n  /*background: #eee; #777*/\n  overflow: hidden;\n}\n/* Center content */\n.mainContent {\n  background: #fff;\n  word-wrap: break-word;\n}\n:host {\n  min-height: 100%;\n}\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9Vc2Vycy93b25kZXIxL2dvL3NyYy9waS9qcy9hcHBzL3NlZWQvc3JjL2FwcC9hcHAuY29tcG9uZW50LnNjc3MiLCJzcmMvYXBwL2FwcC5jb21wb25lbnQuc2NzcyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiQUFDQSxZQUFBO0FBQ0E7RUFDQSx5QkFBQTtFQUNBLGdCQUFBO0FDQUE7QURHQSxtQkFBQTtBQUNBO0VBQ0UsZ0JBQUE7RUFHQSxxQkFBQTtBQ0ZGO0FES0E7RUFDRSxnQkFBQTtBQ0ZGIiwiZmlsZSI6InNyYy9hcHAvYXBwLmNvbXBvbmVudC5zY3NzIiwic291cmNlc0NvbnRlbnQiOlsiXG4vKiBzaWRlYmFyICovXG4uc2lkZWJhciB7XG4vKmJhY2tncm91bmQ6ICNlZWU7ICM3NzcqL1xub3ZlcmZsb3c6IGhpZGRlbjtcbn1cblxuLyogQ2VudGVyIGNvbnRlbnQgKi9cbi5tYWluQ29udGVudCB7XG4gIGJhY2tncm91bmQ6ICNmZmYgLyo5OTkqLztcbiAgLy8gb3ZlcmZsb3cteTogc2Nyb2xsO1xuICAvLyBvdmVyZmxvdzogaGlkZGVuO1xuICB3b3JkLXdyYXA6IGJyZWFrLXdvcmQ7XG59XG5cbjpob3N0IHsgXG4gIG1pbi1oZWlnaHQ6IDEwMCU7XG4gIC8vIGRpc3BsYXk6IGZsZXg7XG59IiwiLyogc2lkZWJhciAqL1xuLnNpZGViYXIge1xuICAvKmJhY2tncm91bmQ6ICNlZWU7ICM3NzcqL1xuICBvdmVyZmxvdzogaGlkZGVuO1xufVxuXG4vKiBDZW50ZXIgY29udGVudCAqL1xuLm1haW5Db250ZW50IHtcbiAgYmFja2dyb3VuZDogI2ZmZjtcbiAgd29yZC13cmFwOiBicmVhay13b3JkO1xufVxuXG46aG9zdCB7XG4gIG1pbi1oZWlnaHQ6IDEwMCU7XG59Il19 */");

/***/ }),

/***/ "./src/app/app.component.ts":
/*!**********************************!*\
  !*** ./src/app/app.component.ts ***!
  \**********************************/
/*! exports provided: AppComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppComponent", function() { return AppComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");


var AppComponent = /** @class */ (function () {
    function AppComponent() {
        this.title = 'seed';
        this.host = window.location.host;
        this.mxNavStyle = {
            'width': '200px',
            'background': '#eee' /*'#efefef'*/
        };
    }
    AppComponent = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'seed',
            template: "\n  <div [ngStyle]=\"{display: 'block', width: '100%'}\">Seed for NAAS @{{host}}</div>\n  <div>\n  <mat-tab-group [ngStyle]=\"{width: '100%'}\">\n    <mat-tab label=\"Pins\"> \n      <pins></pins>\n    </mat-tab>\n    <mat-tab label=\"Field\"> \n      <field-com></field-com>\n      </mat-tab>\n    <mat-tab label=\"Device\"> \n      <device></device>\n    </mat-tab>\n    <mat-tab label=\"Camera\">\n      <camera></camera>\n    </mat-tab>\n  </mat-tab-group>\n  </div>\n  ",
            styles: [Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__importDefault"])(__webpack_require__(/*! ./app.component.scss */ "./src/app/app.component.scss")).default]
        })
    ], AppComponent);
    return AppComponent;
}());



/***/ }),

/***/ "./src/app/app.module.ts":
/*!*******************************!*\
  !*** ./src/app/app.module.ts ***!
  \*******************************/
/*! exports provided: AppModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppModule", function() { return AppModule; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/platform-browser */ "../../node_modules/@angular/platform-browser/__ivy_ngcc__/fesm5/platform-browser.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/common/http */ "../../node_modules/@angular/common/__ivy_ngcc__/fesm5/http.js");
/* harmony import */ var _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/platform-browser/animations */ "../../node_modules/@angular/platform-browser/__ivy_ngcc__/fesm5/animations.js");
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/common */ "../../node_modules/@angular/common/__ivy_ngcc__/fesm5/common.js");
/* harmony import */ var _ngrx_store__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @ngrx/store */ "../../node_modules/@ngrx/store/__ivy_ngcc__/fesm5/store.js");
/* harmony import */ var _ngrx_effects__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @ngrx/effects */ "../../node_modules/@ngrx/effects/__ivy_ngcc__/fesm5/effects.js");
/* harmony import */ var _ngrx_router_store__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @ngrx/router-store */ "../../node_modules/@ngrx/router-store/__ivy_ngcc__/fesm5/router-store.js");
/* harmony import */ var _ngrx_store_devtools__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @ngrx/store-devtools */ "../../node_modules/@ngrx/store-devtools/__ivy_ngcc__/fesm5/store-devtools.js");
/* harmony import */ var ngx_cookie_service__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ngx-cookie-service */ "../../node_modules/ngx-cookie-service/__ivy_ngcc__/ngx-cookie-service.es5.js");
/* harmony import */ var _store__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ./store */ "./src/app/store/index.ts");
/* harmony import */ var _store_effects__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ./store/effects */ "./src/app/store/effects.ts");
/* harmony import */ var _utils_app_routing_module__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ./utils/app-routing.module */ "./src/app/utils/app-routing.module.ts");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @angular/forms */ "../../node_modules/@angular/forms/__ivy_ngcc__/fesm5/forms.js");
/* harmony import */ var _angular_flex_layout__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! @angular/flex-layout */ "../../node_modules/@angular/flex-layout/__ivy_ngcc__/esm5/flex-layout.es5.js");
/* harmony import */ var _utils_material_module__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ./utils/material.module */ "./src/app/utils/material.module.ts");
/* harmony import */ var _utils_CustomRouterStateSerializer__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ./utils/CustomRouterStateSerializer */ "./src/app/utils/CustomRouterStateSerializer.ts");
/* harmony import */ var _auth_interceptor__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ./auth.interceptor */ "./src/app/auth.interceptor.ts");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ../environments/environment */ "./src/environments/environment.ts");
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ./app.component */ "./src/app/app.component.ts");
/* harmony import */ var _containers_field_field_component__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! ./containers/field/field.component */ "./src/app/containers/field/field.component.ts");
/* harmony import */ var _containers_device_device_component__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ./containers/device/device.component */ "./src/app/containers/device/device.component.ts");
/* harmony import */ var _containers_camera_camera_component__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ./containers/camera/camera.component */ "./src/app/containers/camera/camera.component.ts");
/* harmony import */ var _services_naas_service__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! ./services/naas.service */ "./src/app/services/naas.service.ts");
/* harmony import */ var _containers_pins_pins_component__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! ./containers/pins/pins.component */ "./src/app/containers/pins/pins.component.ts");








// import { DBModule } from '@ngrx/db';



















var AppModule = /** @class */ (function () {
    function AppModule() {
    }
    AppModule = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_2__["NgModule"])({
            declarations: [
                _app_component__WEBPACK_IMPORTED_MODULE_20__["AppComponent"],
                _containers_field_field_component__WEBPACK_IMPORTED_MODULE_21__["FieldComponent"], _containers_device_device_component__WEBPACK_IMPORTED_MODULE_22__["DeviceComponent"], _containers_camera_camera_component__WEBPACK_IMPORTED_MODULE_23__["CameraComponent"], _containers_pins_pins_component__WEBPACK_IMPORTED_MODULE_25__["PinsComponent"]
            ],
            imports: [
                _angular_platform_browser__WEBPACK_IMPORTED_MODULE_1__["BrowserModule"],
                _angular_common__WEBPACK_IMPORTED_MODULE_5__["CommonModule"],
                _angular_platform_browser__WEBPACK_IMPORTED_MODULE_1__["BrowserModule"],
                _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_4__["BrowserAnimationsModule"],
                _angular_common_http__WEBPACK_IMPORTED_MODULE_3__["HttpClientModule"],
                /**
                 * StoreModule.forRoot is imported once in the root module, accepting a reducer
                 * function or object map of reducer functions. If passed an object of
                 * reducers, combineReducers will be run creating your application
                 * meta-reducer. This returns all providers for an @ngrx/store
                 * based application.
                 */
                _ngrx_store__WEBPACK_IMPORTED_MODULE_6__["StoreModule"].forRoot(_store__WEBPACK_IMPORTED_MODULE_11__["reducers"], {
                    metaReducers: _store__WEBPACK_IMPORTED_MODULE_11__["metaReducers"]
                }),
                /**
                 * @ngrx/router-store keeps router state up-to-date in the store.
                 */
                _ngrx_router_store__WEBPACK_IMPORTED_MODULE_8__["StoreRouterConnectingModule"].forRoot(),
                /**
                 * Store devtools instrument the store retaining past versions of state
                 * and recalculating new states. This enables powerful time-travel
                 * debugging.
                 *
                 * To use the debugger, install the Redux Devtools extension for either
                 * Chrome or Firefox
                 *
                 * See: https://github.com/zalmoxisus/redux-devtools-extension
                 */
                !_environments_environment__WEBPACK_IMPORTED_MODULE_19__["environment"].production ?
                    _ngrx_store_devtools__WEBPACK_IMPORTED_MODULE_9__["StoreDevtoolsModule"].instrument({
                    // name: 'NgRx Book Store DevTools',
                    }) :
                    [],
                /**
                 * EffectsModule.forRoot() is imported once in the root module and
                 * sets up the effects class to be initialized immediately when the
                 * application starts.
                 *
                 * See: https://github.com/ngrx/platform/blob/master/docs/effects/api.md#forroot
                 */
                _ngrx_effects__WEBPACK_IMPORTED_MODULE_7__["EffectsModule"].forRoot(_store_effects__WEBPACK_IMPORTED_MODULE_12__["CORE_EFFECTS"]),
                /* RouterModule.forRoot([], {
                  initialNavigation: 'enabled'
                }), */
                _utils_app_routing_module__WEBPACK_IMPORTED_MODULE_13__["AppRoutingModule"],
                _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_4__["NoopAnimationsModule"],
                _angular_forms__WEBPACK_IMPORTED_MODULE_14__["FormsModule"],
                _angular_forms__WEBPACK_IMPORTED_MODULE_14__["ReactiveFormsModule"],
                _angular_flex_layout__WEBPACK_IMPORTED_MODULE_15__["FlexLayoutModule"],
                _utils_material_module__WEBPACK_IMPORTED_MODULE_16__["MaterialModule"],
            ],
            providers: [
                /**
                 * The `RouterStateSnapshot` provided by the `Router` is a large complex structure.
                 * A custom RouterStateSerializer is used to parse the `RouterStateSnapshot` provided
                 * by `@ngrx/router-store` to include only the desired pieces of the snapshot.
                 */
                { provide: _ngrx_router_store__WEBPACK_IMPORTED_MODULE_8__["RouterStateSerializer"], useClass: _utils_CustomRouterStateSerializer__WEBPACK_IMPORTED_MODULE_17__["CustomRouterStateSerializer"] },
                { provide: _angular_common_http__WEBPACK_IMPORTED_MODULE_3__["HTTP_INTERCEPTORS"], useClass: _auth_interceptor__WEBPACK_IMPORTED_MODULE_18__["AuthInterceptor"], multi: true },
                ngx_cookie_service__WEBPACK_IMPORTED_MODULE_10__["CookieService"],
                _services_naas_service__WEBPACK_IMPORTED_MODULE_24__["NaasService"]
            ],
            bootstrap: [_app_component__WEBPACK_IMPORTED_MODULE_20__["AppComponent"]]
        })
    ], AppModule);
    return AppModule;
}());



/***/ }),

/***/ "./src/app/auth.interceptor.ts":
/*!*************************************!*\
  !*** ./src/app/auth.interceptor.ts ***!
  \*************************************/
/*! exports provided: AuthInterceptor */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AuthInterceptor", function() { return AuthInterceptor; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");


var AuthInterceptor = /** @class */ (function () {
    function AuthInterceptor() {
    }
    AuthInterceptor.prototype.intercept = function (req, next) {
        // const authReq = req.clone({
        //   setHeaders: { Authorization: `Bearer authtest` }
        // });
        // console.log('intercept');
        console.log(req);
        // console.log(next);
        return next.handle(req);
    };
    AuthInterceptor = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])()
    ], AuthInterceptor);
    return AuthInterceptor;
}());



/***/ }),

/***/ "./src/app/containers/camera/camera.component.scss":
/*!*********************************************************!*\
  !*** ./src/app/containers/camera/camera.component.scss ***!
  \*********************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ("\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2NvbnRhaW5lcnMvY2FtZXJhL2NhbWVyYS5jb21wb25lbnQuc2NzcyJ9 */");

/***/ }),

/***/ "./src/app/containers/camera/camera.component.ts":
/*!*******************************************************!*\
  !*** ./src/app/containers/camera/camera.component.ts ***!
  \*******************************************************/
/*! exports provided: CameraComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CameraComponent", function() { return CameraComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");


var CameraComponent = /** @class */ (function () {
    function CameraComponent() {
    }
    CameraComponent.prototype.ngOnInit = function () {
    };
    CameraComponent = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'camera',
            template: "./camera.component.html",
            styles: [Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__importDefault"])(__webpack_require__(/*! ./camera.component.scss */ "./src/app/containers/camera/camera.component.scss")).default]
        }),
        Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"])("design:paramtypes", [])
    ], CameraComponent);
    return CameraComponent;
}());



/***/ }),

/***/ "./src/app/containers/device/device.component.scss":
/*!*********************************************************!*\
  !*** ./src/app/containers/device/device.component.scss ***!
  \*********************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ("\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2NvbnRhaW5lcnMvZGV2aWNlL2RldmljZS5jb21wb25lbnQuc2NzcyJ9 */");

/***/ }),

/***/ "./src/app/containers/device/device.component.ts":
/*!*******************************************************!*\
  !*** ./src/app/containers/device/device.component.ts ***!
  \*******************************************************/
/*! exports provided: DeviceComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "DeviceComponent", function() { return DeviceComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");


var DeviceComponent = /** @class */ (function () {
    function DeviceComponent() {
    }
    DeviceComponent.prototype.ngOnInit = function () {
    };
    DeviceComponent = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'device',
            template: "./device.component.html",
            styles: [Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__importDefault"])(__webpack_require__(/*! ./device.component.scss */ "./src/app/containers/device/device.component.scss")).default]
        }),
        Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"])("design:paramtypes", [])
    ], DeviceComponent);
    return DeviceComponent;
}());



/***/ }),

/***/ "./src/app/containers/field/field.component.scss":
/*!*******************************************************!*\
  !*** ./src/app/containers/field/field.component.scss ***!
  \*******************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ("\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2NvbnRhaW5lcnMvZmllbGQvZmllbGQuY29tcG9uZW50LnNjc3MifQ== */");

/***/ }),

/***/ "./src/app/containers/field/field.component.ts":
/*!*****************************************************!*\
  !*** ./src/app/containers/field/field.component.ts ***!
  \*****************************************************/
/*! exports provided: FieldComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "FieldComponent", function() { return FieldComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _ngrx_store__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @ngrx/store */ "../../node_modules/@ngrx/store/__ivy_ngcc__/fesm5/store.js");
/* harmony import */ var _services_naas_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../../services/naas.service */ "./src/app/services/naas.service.ts");




var FieldComponent = /** @class */ (function () {
    /* fieldArray: Dictionary<IFile>;
    fileEntities$ = this.store.select(fileSelector.getFileEntities);
    fileEntities$$ = this.fileEntities$.subscribe((fileEntities: Dictionary<IFile>) => {
      this.fileEntities = fileEntities;
    }); */
    function FieldComponent(store, naas) {
        this.store = store;
        this.naas = naas;
        this.fields = this.naas.getFields();
        this.devices = this.naas.readDevices();
    }
    FieldComponent.prototype.ngOnInit = function () {
    };
    FieldComponent.ctorParameters = function () { return [
        { type: _ngrx_store__WEBPACK_IMPORTED_MODULE_2__["Store"] },
        { type: _services_naas_service__WEBPACK_IMPORTED_MODULE_3__["NaasService"] }
    ]; };
    FieldComponent = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'field-com',
            template: "\n  <mat-accordion *ngIf=\"this.fields\">\n    <mat-expansion-panel *ngFor=\"let field of this.fields | async\">\n      <mat-expansion-panel-header>\n        <mat-panel-title>\n          {{field.Name}}\n        </mat-panel-title>\n        <mat-panel-description>\n        </mat-panel-description>\n      </mat-expansion-panel-header>\n      \n      <div *ngIf=\"this.fields\">Field</div>\n      \n      \n      <div *ngIf=\"this.plants\">Plant</div>\n      \n      \n    </mat-expansion-panel>\n  </mat-accordion>\n  ",
            styles: [Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__importDefault"])(__webpack_require__(/*! ./field.component.scss */ "./src/app/containers/field/field.component.scss")).default]
        }),
        Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"])("design:paramtypes", [_ngrx_store__WEBPACK_IMPORTED_MODULE_2__["Store"], _services_naas_service__WEBPACK_IMPORTED_MODULE_3__["NaasService"]])
    ], FieldComponent);
    return FieldComponent;
}());



/***/ }),

/***/ "./src/app/containers/pins/pins.component.scss":
/*!*****************************************************!*\
  !*** ./src/app/containers/pins/pins.component.scss ***!
  \*****************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ("\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL2NvbnRhaW5lcnMvcGlucy9waW5zLmNvbXBvbmVudC5zY3NzIn0= */");

/***/ }),

/***/ "./src/app/containers/pins/pins.component.ts":
/*!***************************************************!*\
  !*** ./src/app/containers/pins/pins.component.ts ***!
  \***************************************************/
/*! exports provided: PinsComponent */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "PinsComponent", function() { return PinsComponent; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! rxjs */ "../../node_modules/rxjs/_esm5/index.js");
/* harmony import */ var _services_naas_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../../services/naas.service */ "./src/app/services/naas.service.ts");




var PinsComponent = /** @class */ (function () {
    function PinsComponent(naas) {
        this.naas = naas;
        this.pins = [];
        this.pinsState$ = new rxjs__WEBPACK_IMPORTED_MODULE_2__["BehaviorSubject"]([]);
    }
    PinsComponent.prototype.ngOnInit = function () {
        this.pins.push({ Pin: 1, Status: false });
        this.pins.push({ Pin: 2, Status: false });
        this.pins.push({ Pin: 3, Status: false });
        this.pins.push({ Pin: 4, Status: false });
        this.pins.push({ Pin: 5, Status: false });
        // this.pins.push({Pin: 6, Status: false}); 
        this.pins.push({ Pin: 7, Status: false });
        this.pins.push({ Pin: 8, Status: false });
        // this.pins.push({Pin: 9, Status: false});
        this.pins.push({ Pin: 10, Status: false });
        this.pins.push({ Pin: 11, Status: false });
        this.pins.push({ Pin: 12, Status: false });
        this.pins.push({ Pin: 13, Status: false });
        // this.pins.push({Pin: 14, Status: false}); 
        this.pins.push({ Pin: 15, Status: false });
        this.pins.push({ Pin: 16, Status: false });
        this.pins.push({ Pin: 17, Status: false });
        this.pins.push({ Pin: 18, Status: false });
        this.pins.push({ Pin: 19, Status: false });
        // this.pins.push({Pin: 20, Status: false});
        this.pins.push({ Pin: 21, Status: false });
        this.pins.push({ Pin: 22, Status: false });
        this.pins.push({ Pin: 23, Status: false });
        this.pins.push({ Pin: 24, Status: false });
        // this.pins.push({Pin: 25, Status: false});
        this.pins.push({ Pin: 26, Status: false });
        this.pins.push({ Pin: 27, Status: false });
        this.pins.push({ Pin: 28, Status: false });
        this.pins.push({ Pin: 29, Status: false });
        // this.pins.push({Pin: 30, Status: false});
        this.pins.push({ Pin: 31, Status: false });
        this.pins.push({ Pin: 32, Status: false });
        this.pins.push({ Pin: 33, Status: false });
        // this.pins.push({Pin: 34, Status: false});
        this.pins.push({ Pin: 35, Status: false });
        this.pins.push({ Pin: 36, Status: false });
        this.pins.push({ Pin: 37, Status: false });
        this.pins.push({ Pin: 38, Status: false });
        // this.pins.push({Pin: 39, Status: false});
        this.pins.push({ Pin: 40, Status: false });
        this.pinsState$.next(this.pins);
        console.log('pin');
    };
    PinsComponent.prototype.togglePin = function (pinNo) {
        var _this = this;
        if (pinNo > 0) {
            var _loop_1 = function (pin) {
                if (pin.Pin === pinNo) {
                    if (pin.Status) {
                        this_1.naas.pinOff(pin.Pin).subscribe(function (result) {
                            pin.Status = false;
                            _this.emitPins(_this.pins);
                        });
                        return "break";
                    }
                    else {
                        this_1.naas.pinOn(pin.Pin).subscribe(function (result) {
                            pin.Status = true;
                            _this.emitPins(_this.pins);
                        });
                        return "break";
                    }
                }
            };
            var this_1 = this;
            for (var _i = 0, _a = this.pins; _i < _a.length; _i++) {
                var pin = _a[_i];
                var state_1 = _loop_1(pin);
                if (state_1 === "break")
                    break;
            }
        }
    };
    PinsComponent.prototype.emitPins = function (pins) {
        this.pinsState$.next(pins);
    };
    PinsComponent.prototype.handlePinsState = function (pinsState) {
        if (pinsState) {
            console.log("handlePinsState");
            var ps = JSON.parse(pinsState);
            var pins = [];
            // push to pinsState$
            this.emitPins(pins);
        }
    };
    PinsComponent.ctorParameters = function () { return [
        { type: _services_naas_service__WEBPACK_IMPORTED_MODULE_3__["NaasService"] }
    ]; };
    PinsComponent = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
            selector: 'pins',
            template: "\n  <div *ngFor=\"let pin of pinsState$ | async\" >\n    <mat-slide-toggle\n    (change)=\"togglePin(pin.Pin)\"\n    [checked]=\"pin.Status\">{{pin.Pin}}</mat-slide-toggle>\n  </div>\n  ",
            styles: [Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__importDefault"])(__webpack_require__(/*! ./pins.component.scss */ "./src/app/containers/pins/pins.component.scss")).default]
        }),
        Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"])("design:paramtypes", [_services_naas_service__WEBPACK_IMPORTED_MODULE_3__["NaasService"]])
    ], PinsComponent);
    return PinsComponent;
}());



/***/ }),

/***/ "./src/app/services/naas.service.ts":
/*!******************************************!*\
  !*** ./src/app/services/naas.service.ts ***!
  \******************************************/
/*! exports provided: NaasService */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "NaasService", function() { return NaasService; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common/http */ "../../node_modules/@angular/common/__ivy_ngcc__/fesm5/http.js");
/* harmony import */ var _ngrx_store__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @ngrx/store */ "../../node_modules/@ngrx/store/__ivy_ngcc__/fesm5/store.js");
/* harmony import */ var _store_field_field_selector__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../store/field/field.selector */ "./src/app/store/field/field.selector.ts");





var NaasService = /** @class */ (function () {
    function NaasService(http, store) {
        var _this = this;
        this.http = http;
        this.store = store;
        this.host = window.location.hostname;
        this.prefix = "http://" + this.host + ":3030/api/";
        this.fieldPrefix = this.prefix + "field/";
        this.devicePrefix = this.prefix + "device/";
        this.httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpHeaders"]({
                'Content-Type': 'application/json',
            })
        };
        this.fieldArray$ = this.store.select(_store_field_field_selector__WEBPACK_IMPORTED_MODULE_4__["getAllFields"]);
        this.fieldArray$$ = this.fieldArray$.subscribe(function (fieldArray) {
            _this.fieldArray = fieldArray;
        });
    }
    NaasService.prototype.getFields = function () {
        return this.fieldArray$;
    };
    NaasService.prototype.readFields = function () {
        console.log('readFields');
        return this.http.post(this.fieldPrefix + 'list', {}, this.httpOptions);
    };
    NaasService.prototype.readDevices = function () {
        console.log('readDevices');
        return this.http.post(this.devicePrefix + 'list', {}, this.httpOptions);
    };
    NaasService.prototype.pinOn = function (pin) {
        if (pin) {
            console.log('pinOn');
            return this.http.post(this.devicePrefix + 'on/' + pin, {}, this.httpOptions);
        }
    };
    NaasService.prototype.pinOff = function (pin) {
        if (pin) {
            console.log('pinOff');
            return this.http.post(this.devicePrefix + 'off/' + pin, {}, this.httpOptions);
        }
    };
    NaasService.ctorParameters = function () { return [
        { type: _angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpClient"] },
        { type: _ngrx_store__WEBPACK_IMPORTED_MODULE_3__["Store"] }
    ]; };
    NaasService = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])({
            providedIn: 'root'
        }),
        Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"])("design:paramtypes", [_angular_common_http__WEBPACK_IMPORTED_MODULE_2__["HttpClient"], _ngrx_store__WEBPACK_IMPORTED_MODULE_3__["Store"]])
    ], NaasService);
    return NaasService;
}());



/***/ }),

/***/ "./src/app/store/device/device.action.ts":
/*!***********************************************!*\
  !*** ./src/app/store/device/device.action.ts ***!
  \***********************************************/
/*! exports provided: DeviceActionTypes, Y */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "DeviceActionTypes", function() { return DeviceActionTypes; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "Y", function() { return Y; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");

var DeviceActionTypes;
(function (DeviceActionTypes) {
    DeviceActionTypes["Y"] = "[device] Y";
})(DeviceActionTypes || (DeviceActionTypes = {}));
var Y = /** @class */ (function () {
    function Y() {
        this.type = DeviceActionTypes.Y;
    }
    return Y;
}());



/***/ }),

/***/ "./src/app/store/device/device.reducer.ts":
/*!************************************************!*\
  !*** ./src/app/store/device/device.reducer.ts ***!
  \************************************************/
/*! exports provided: adapter, initialState, reducer, getY */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "adapter", function() { return adapter; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "initialState", function() { return initialState; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "reducer", function() { return reducer; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getY", function() { return getY; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _device_action__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./device.action */ "./src/app/store/device/device.action.ts");
/* harmony import */ var _ngrx_entity__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @ngrx/entity */ "../../node_modules/@ngrx/entity/__ivy_ngcc__/fesm5/entity.js");



var adapter = Object(_ngrx_entity__WEBPACK_IMPORTED_MODULE_2__["createEntityAdapter"])({
    selectId: function (device) { return device.Id; },
    sortComparer: false,
});
var initialState = adapter.getInitialState({
    y: null
});
/* export interface State {
    y: boolean;
}

const initialState: State = {
    y: true,
}; */
function reducer(state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case _device_action__WEBPACK_IMPORTED_MODULE_1__["DeviceActionTypes"].Y:
            return Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__assign"])({}, state);
        default:
            return state;
    }
}
var getY = function (state) { return state.y; };


/***/ }),

/***/ "./src/app/store/effects.ts":
/*!**********************************!*\
  !*** ./src/app/store/effects.ts ***!
  \**********************************/
/*! exports provided: CORE_EFFECTS */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CORE_EFFECTS", function() { return CORE_EFFECTS; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
// import { MxNavEffects } from './mx-nav/mx-nav.effect';

var CORE_EFFECTS = [
// MxNavEffects,
];


/***/ }),

/***/ "./src/app/store/field/field.action.ts":
/*!*********************************************!*\
  !*** ./src/app/store/field/field.action.ts ***!
  \*********************************************/
/*! exports provided: FieldActionTypes, Y */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "FieldActionTypes", function() { return FieldActionTypes; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "Y", function() { return Y; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");

var FieldActionTypes;
(function (FieldActionTypes) {
    FieldActionTypes["Y"] = "[field] Y";
})(FieldActionTypes || (FieldActionTypes = {}));
var Y = /** @class */ (function () {
    function Y() {
        this.type = FieldActionTypes.Y;
    }
    return Y;
}());



/***/ }),

/***/ "./src/app/store/field/field.reducer.ts":
/*!**********************************************!*\
  !*** ./src/app/store/field/field.reducer.ts ***!
  \**********************************************/
/*! exports provided: adapter, initialState, reducer */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "adapter", function() { return adapter; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "initialState", function() { return initialState; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "reducer", function() { return reducer; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _ngrx_entity__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @ngrx/entity */ "../../node_modules/@ngrx/entity/__ivy_ngcc__/fesm5/entity.js");
/* harmony import */ var _field_action__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./field.action */ "./src/app/store/field/field.action.ts");



var adapter = Object(_ngrx_entity__WEBPACK_IMPORTED_MODULE_1__["createEntityAdapter"])({
    selectId: function (field) { return field.Id; },
    sortComparer: false,
});
var initialState = adapter.getInitialState({
// StarterResource: null,
});
/* export interface State {
    y: boolean;
}

const initialState: State = {
    y: true,
}; */
function reducer(state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case _field_action__WEBPACK_IMPORTED_MODULE_2__["FieldActionTypes"].Y:
            return Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__assign"])({}, state);
        default:
            return state;
    }
}
// export const getY = (state: State) => state.y;


/***/ }),

/***/ "./src/app/store/field/field.selector.ts":
/*!***********************************************!*\
  !*** ./src/app/store/field/field.selector.ts ***!
  \***********************************************/
/*! exports provided: getFieldIds, getFieldEntities, getAllFields, getFieldTotal */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getFieldIds", function() { return getFieldIds; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getFieldEntities", function() { return getFieldEntities; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getAllFields", function() { return getAllFields; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getFieldTotal", function() { return getFieldTotal; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _index__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./../index */ "./src/app/store/index.ts");
/* harmony import */ var _field_reducer__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./field.reducer */ "./src/app/store/field/field.reducer.ts");
var _a;

// from Feature


/* export const getFieldState = createSelector(
  fromRoot.getFieldState,
  (state: fromRoot.State) => state.field
); */
// export const getSyntaxEntitiesState = createSelector(
//   fromDeskIndex.getDeskState,
//   state => state.syntax
// );
var 
// select the array of ids
getFieldIds = (_a = _field_reducer__WEBPACK_IMPORTED_MODULE_2__["adapter"].getSelectors(_index__WEBPACK_IMPORTED_MODULE_1__["getFieldState"]), _a.selectIds), 
// select the dictionary of entities
getFieldEntities = _a.selectEntities, 
// select the array of the entity
getAllFields = _a.selectAll, 
// select the total count
getFieldTotal = _a.selectTotal;
/*
export const getXState = createSelector(
  fromX.getY,
  (state: fromF.FState) => state.x
);

export const getY = createSelector(
  getXState,
  fromX.getY
); */


/***/ }),

/***/ "./src/app/store/index.ts":
/*!********************************!*\
  !*** ./src/app/store/index.ts ***!
  \********************************/
/*! exports provided: reducers, logger, metaReducers, getFieldState, getPlantState, getDeviceState */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "reducers", function() { return reducers; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "logger", function() { return logger; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "metaReducers", function() { return metaReducers; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getFieldState", function() { return getFieldState; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getPlantState", function() { return getPlantState; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getDeviceState", function() { return getDeviceState; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../../environments/environment */ "./src/environments/environment.ts");
/* harmony import */ var _ngrx_router_store__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @ngrx/router-store */ "../../node_modules/@ngrx/router-store/__ivy_ngcc__/fesm5/router-store.js");
/* harmony import */ var ngrx_store_freeze__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ngrx-store-freeze */ "../../node_modules/ngrx-store-freeze/index.js");
/* harmony import */ var ngrx_store_freeze__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(ngrx_store_freeze__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _plant_plant_reducer__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./plant/plant.reducer */ "./src/app/store/plant/plant.reducer.ts");
/* harmony import */ var _device_device_reducer__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./device/device.reducer */ "./src/app/store/device/device.reducer.ts");
/* harmony import */ var _field_field_reducer__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ./field/field.reducer */ "./src/app/store/field/field.reducer.ts");



/**
 * storeFreeze prevents state from being mutated. When mutation occurs, an
 * exception will be thrown. This is useful during development mode to
 * ensure that none of the reducers accidentally mutates the state.
 */

/**
 * Every reducer module's default export is the reducer function itself. In
 * addition, each module should export a type or interface that describes
 * the state of the reducer plus any selector functions. The `* as`
 * notation packages up all of the exports into a single object.
 */



/**
 * Our state is composed of a map of action reducer functions.
 * These reducer functions are called with each dispatched action
 * and the current or initial state and return a new immutable state.
 */
var reducers = {
    plant: _plant_plant_reducer__WEBPACK_IMPORTED_MODULE_4__["reducer"],
    device: _device_device_reducer__WEBPACK_IMPORTED_MODULE_5__["reducer"],
    field: _field_field_reducer__WEBPACK_IMPORTED_MODULE_6__["reducer"],
    routerReducer: _ngrx_router_store__WEBPACK_IMPORTED_MODULE_2__["routerReducer"],
};
// console.log all actions
function logger(reducer) {
    return function (state, action) {
        console.log('state', state);
        console.log('action', action);
        return reducer(state, action);
    };
}
/**
 * By default, @ngrx/store uses combineReducers with the reducer map to compose
 * the root meta-reducer. To add more meta-reducers, provide an array of meta-reducers
 * that will be composed to form the root meta-reducer.
 */
var metaReducers = !_environments_environment__WEBPACK_IMPORTED_MODULE_1__["environment"].production
    ? [logger, ngrx_store_freeze__WEBPACK_IMPORTED_MODULE_3__["storeFreeze"]] // , storeFreeze
    : [];
var getFieldState = function (state) { return state.field; };
var getPlantState = function (state) { return state.plant; };
var getDeviceState = function (state) { return state.device; };


/***/ }),

/***/ "./src/app/store/plant/plant.action.ts":
/*!*********************************************!*\
  !*** ./src/app/store/plant/plant.action.ts ***!
  \*********************************************/
/*! exports provided: PlantActionTypes, Y */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "PlantActionTypes", function() { return PlantActionTypes; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "Y", function() { return Y; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");

var PlantActionTypes;
(function (PlantActionTypes) {
    PlantActionTypes["Y"] = "[plant] Y";
})(PlantActionTypes || (PlantActionTypes = {}));
var Y = /** @class */ (function () {
    function Y() {
        this.type = PlantActionTypes.Y;
    }
    return Y;
}());



/***/ }),

/***/ "./src/app/store/plant/plant.reducer.ts":
/*!**********************************************!*\
  !*** ./src/app/store/plant/plant.reducer.ts ***!
  \**********************************************/
/*! exports provided: reducer, getY */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "reducer", function() { return reducer; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getY", function() { return getY; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _plant_action__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./plant.action */ "./src/app/store/plant/plant.action.ts");


var initialState = {
    y: true,
};
function reducer(state, action) {
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case _plant_action__WEBPACK_IMPORTED_MODULE_1__["PlantActionTypes"].Y:
            return Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__assign"])(Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__assign"])({}, state), { y: true });
        default:
            return state;
    }
}
var getY = function (state) { return state.y; };


/***/ }),

/***/ "./src/app/utils/CustomRouterStateSerializer.ts":
/*!******************************************************!*\
  !*** ./src/app/utils/CustomRouterStateSerializer.ts ***!
  \******************************************************/
/*! exports provided: CustomRouterStateSerializer */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "CustomRouterStateSerializer", function() { return CustomRouterStateSerializer; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");


var CustomRouterStateSerializer = /** @class */ (function () {
    function CustomRouterStateSerializer() {
    }
    CustomRouterStateSerializer.prototype.serialize = function (routerState) {
        var url = routerState.url;
        var queryParams = routerState.root.queryParams;
        return { url: url, queryParams: queryParams };
    };
    CustomRouterStateSerializer = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Injectable"])()
    ], CustomRouterStateSerializer);
    return CustomRouterStateSerializer;
}());



/***/ }),

/***/ "./src/app/utils/app-routing.module.ts":
/*!*********************************************!*\
  !*** ./src/app/utils/app-routing.module.ts ***!
  \*********************************************/
/*! exports provided: AppRoutingModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "AppRoutingModule", function() { return AppRoutingModule; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/router */ "../../node_modules/@angular/router/__ivy_ngcc__/fesm5/router.js");
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./../app.component */ "./src/app/app.component.ts");




// import { PageNotFoundComponent } from './containers/page-not-found/page-not-found.component';
var routes = [
    { path: '', component: _app_component__WEBPACK_IMPORTED_MODULE_3__["AppComponent"] },
    { path: 'abc', component: _app_component__WEBPACK_IMPORTED_MODULE_3__["AppComponent"] },
];
var AppRoutingModule = /** @class */ (function () {
    function AppRoutingModule() {
    }
    AppRoutingModule = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["NgModule"])({
            imports: [_angular_router__WEBPACK_IMPORTED_MODULE_2__["RouterModule"].forRoot(routes)],
            exports: [_angular_router__WEBPACK_IMPORTED_MODULE_2__["RouterModule"]]
        })
    ], AppRoutingModule);
    return AppRoutingModule;
}());



/***/ }),

/***/ "./src/app/utils/material.module.ts":
/*!******************************************!*\
  !*** ./src/app/utils/material.module.ts ***!
  \******************************************/
/*! exports provided: MATERIALS, MaterialModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "MATERIALS", function() { return MATERIALS; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "MaterialModule", function() { return MaterialModule; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _angular_material_autocomplete__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/material/autocomplete */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/autocomplete.js");
/* harmony import */ var _angular_material_button__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/material/button */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/button.js");
/* harmony import */ var _angular_material_button_toggle__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/material/button-toggle */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/button-toggle.js");
/* harmony import */ var _angular_material_card__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/material/card */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/card.js");
/* harmony import */ var _angular_material_checkbox__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/material/checkbox */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/checkbox.js");
/* harmony import */ var _angular_material_dialog__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @angular/material/dialog */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/dialog.js");
/* harmony import */ var _angular_material_expansion__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @angular/material/expansion */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/expansion.js");
/* harmony import */ var _angular_material_form_field__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @angular/material/form-field */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/form-field.js");
/* harmony import */ var _angular_material_grid_list__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @angular/material/grid-list */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/grid-list.js");
/* harmony import */ var _angular_material_icon__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! @angular/material/icon */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/icon.js");
/* harmony import */ var _angular_material_input__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! @angular/material/input */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/input.js");
/* harmony import */ var _angular_material_list__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! @angular/material/list */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/list.js");
/* harmony import */ var _angular_material_menu__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @angular/material/menu */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/menu.js");
/* harmony import */ var _angular_material_radio__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! @angular/material/radio */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/radio.js");
/* harmony import */ var _angular_material_select__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! @angular/material/select */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/select.js");
/* harmony import */ var _angular_material_sidenav__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @angular/material/sidenav */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/sidenav.js");
/* harmony import */ var _angular_material_slide_toggle__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! @angular/material/slide-toggle */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/slide-toggle.js");
/* harmony import */ var _angular_material_tabs__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! @angular/material/tabs */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/tabs.js");
/* harmony import */ var _angular_material_tooltip__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! @angular/material/tooltip */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/tooltip.js");
/* harmony import */ var _angular_material_tree__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! @angular/material/tree */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/tree.js");
/* harmony import */ var _angular_material_chips__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! @angular/material/chips */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/chips.js");
/* harmony import */ var _angular_material_divider__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! @angular/material/divider */ "../../node_modules/@angular/material/__ivy_ngcc__/fesm5/divider.js");
























var MATERIALS = [
    _angular_material_tooltip__WEBPACK_IMPORTED_MODULE_20__["MatTooltipModule"],
    _angular_material_sidenav__WEBPACK_IMPORTED_MODULE_17__["MatSidenavModule"],
    _angular_material_button__WEBPACK_IMPORTED_MODULE_3__["MatButtonModule"],
    _angular_material_grid_list__WEBPACK_IMPORTED_MODULE_10__["MatGridListModule"],
    _angular_material_select__WEBPACK_IMPORTED_MODULE_16__["MatSelectModule"],
    _angular_material_menu__WEBPACK_IMPORTED_MODULE_14__["MatMenuModule"],
    _angular_material_input__WEBPACK_IMPORTED_MODULE_12__["MatInputModule"],
    _angular_material_radio__WEBPACK_IMPORTED_MODULE_15__["MatRadioModule"],
    _angular_material_list__WEBPACK_IMPORTED_MODULE_13__["MatListModule"],
    _angular_material_form_field__WEBPACK_IMPORTED_MODULE_9__["MatFormFieldModule"],
    _angular_material_card__WEBPACK_IMPORTED_MODULE_5__["MatCardModule"],
    _angular_material_checkbox__WEBPACK_IMPORTED_MODULE_6__["MatCheckboxModule"],
    _angular_material_tree__WEBPACK_IMPORTED_MODULE_21__["MatTreeModule"],
    _angular_material_icon__WEBPACK_IMPORTED_MODULE_11__["MatIconModule"],
    _angular_material_tabs__WEBPACK_IMPORTED_MODULE_19__["MatTabsModule"],
    _angular_material_expansion__WEBPACK_IMPORTED_MODULE_8__["MatExpansionModule"],
    _angular_material_dialog__WEBPACK_IMPORTED_MODULE_7__["MatDialogModule"],
    _angular_material_autocomplete__WEBPACK_IMPORTED_MODULE_2__["MatAutocompleteModule"],
    _angular_material_slide_toggle__WEBPACK_IMPORTED_MODULE_18__["MatSlideToggleModule"],
    _angular_material_button_toggle__WEBPACK_IMPORTED_MODULE_4__["MatButtonToggleModule"],
    _angular_material_chips__WEBPACK_IMPORTED_MODULE_22__["MatChipsModule"],
    _angular_material_divider__WEBPACK_IMPORTED_MODULE_23__["MatDividerModule"],
];
var MaterialModule = /** @class */ (function () {
    function MaterialModule() {
    }
    MaterialModule = Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"])([
        Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["NgModule"])({
            imports: Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__spreadArrays"])(MATERIALS),
            exports: Object(tslib__WEBPACK_IMPORTED_MODULE_0__["__spreadArrays"])(MATERIALS),
        })
    ], MaterialModule);
    return MaterialModule;
}());



/***/ }),

/***/ "./src/environments/environment.ts":
/*!*****************************************!*\
  !*** ./src/environments/environment.ts ***!
  \*****************************************/
/*! exports provided: environment */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "environment", function() { return environment; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

var environment = {
    production: false
};
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.


/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/*! no exports provided */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "../../node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "../../node_modules/@angular/core/__ivy_ngcc__/fesm5/core.js");
/* harmony import */ var _angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/platform-browser-dynamic */ "../../node_modules/@angular/platform-browser-dynamic/__ivy_ngcc__/fesm5/platform-browser-dynamic.js");
/* harmony import */ var _app_app_module__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./app/app.module */ "./src/app/app.module.ts");
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./environments/environment */ "./src/environments/environment.ts");





if (_environments_environment__WEBPACK_IMPORTED_MODULE_4__["environment"].production) {
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["enableProdMode"])();
}
Object(_angular_platform_browser_dynamic__WEBPACK_IMPORTED_MODULE_2__["platformBrowserDynamic"])()
    .bootstrapModule(_app_app_module__WEBPACK_IMPORTED_MODULE_3__["AppModule"])
    .catch(function (err) { return console.error(err); });


/***/ }),

/***/ 0:
/*!***************************!*\
  !*** multi ./src/main.ts ***!
  \***************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(/*! /Users/wonder1/go/src/pi/js/apps/seed/src/main.ts */"./src/main.ts");


/***/ })

},[[0,"runtime","vendor"]]]);
//# sourceMappingURL=main.js.map