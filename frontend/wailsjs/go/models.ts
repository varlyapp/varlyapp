/* Do not change, this code is generated from Golang structs */

export {};

export class NewCollectionConfig {


    static createFrom(source: any = {}) {
        return new NewCollectionConfig(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Document {
    title: string;
    path: string;
    preview: string;

    static createFrom(source: any = {}) {
        return new Document(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.title = source["title"];
        this.path = source["path"];
        this.preview = source["preview"];
    }
}
export class Settings {
    path: string;
    theme: string;
    documents: Document[];

    static createFrom(source: any = {}) {
        return new Settings(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.path = source["path"];
        this.theme = source["theme"];
        this.documents = this.convertValues(source["documents"], Document);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
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
export class MessageDialogOptions {


    static createFrom(source: any = {}) {
        return new MessageDialogOptions(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class CollectionConfig {


    static createFrom(source: any = {}) {
        return new CollectionConfig(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}

