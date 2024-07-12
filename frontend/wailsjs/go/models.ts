export namespace main {
	
	export class ColumnInfo {
	    name: string;
	    type: string;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new ColumnInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.size = source["size"];
	    }
	}
	export class TablelInfo {
	    table_name: string;
	
	    static createFrom(source: any = {}) {
	        return new TablelInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.table_name = source["table_name"];
	    }
	}

}

export namespace model {
	
	export class CodeTemplate {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    fileName: string;
	    fileType: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new CodeTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.fileType = source["fileType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DataModel {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    makeType: number;
	    tableName: string;
	    tpIds: string[];
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new DataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.tpIds = source["tpIds"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DataSource {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new DataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ModelInfo {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    dataModelId: string;
	    title: string;
	    name: string;
	    columnType: string;
	    size: number;
	    isPk: number;
	    extension: string[];
	    controlType: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new ModelInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.dataModelId = source["dataModelId"];
	        this.title = source["title"];
	        this.name = source["name"];
	        this.columnType = source["columnType"];
	        this.size = source["size"];
	        this.isPk = source["isPk"];
	        this.extension = source["extension"];
	        this.controlType = source["controlType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace repository {
	
	export class CreateDataSource {
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateDataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	}
	export class RequestCreateDataModel {
	    title: string;
	    makeType: number;
	    tableName: string;
	    tpIds: string[];
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestCreateDataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.tpIds = source["tpIds"];
	        this.mark = source["mark"];
	    }
	}
	export class RequestCreateTemplate {
	    title: string;
	    fileName: string;
	    fileType: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestCreateTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.fileType = source["fileType"];
	        this.mark = source["mark"];
	    }
	}
	export class RequestModelInfoSave {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    dataModelId: string;
	    title: string;
	    name: string;
	    columnType: string;
	    size: number;
	    isPk: number;
	    extension: string[];
	    controlType: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestModelInfoSave(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.dataModelId = source["dataModelId"];
	        this.title = source["title"];
	        this.name = source["name"];
	        this.columnType = source["columnType"];
	        this.size = source["size"];
	        this.isPk = source["isPk"];
	        this.extension = source["extension"];
	        this.controlType = source["controlType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RequestModifyDataModel {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    makeType: number;
	    tableName: string;
	    tpIds: string[];
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestModifyDataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.tpIds = source["tpIds"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RequestModifyTemplate {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    fileName: string;
	    fileType: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestModifyTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.fileType = source["fileType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RequestTemplateQuery {
	    condition: string;
	    ids: string[];
	
	    static createFrom(source: any = {}) {
	        return new RequestTemplateQuery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.condition = source["condition"];
	        this.ids = source["ids"];
	    }
	}
	export class UpdateDataSource {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateDataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

