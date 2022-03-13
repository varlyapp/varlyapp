export interface go {
  "main": {
    "App": {
		GenerateCollection(arg1:string):Promise<CollectionConfig>
		GenerateCollectionFromConfig(arg1:NewCollectionConfig):Promise<void>
		OpenDirectoryDialog():Promise<string>
		OpenFileDialog():Promise<string>
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
