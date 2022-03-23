export interface go {
  "main": {
    "App": {
		EncodeImage(arg1:string):Promise<string>
		GenerateCollectionFromConfig(arg1:NewCollectionConfig):Promise<void>
		GetApplicationDocumentsDirectory():Promise<string>
		OpenDirectoryDialog():Promise<string>
		OpenFileDialog():Promise<string>
		ReadLayers(arg1:string):Promise<CollectionConfig>
		Save(arg1:string,arg2:string):Promise<boolean>
		SaveFileDialog():Promise<string>
		UploadCollection():Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
