import * as models from './models';

export interface go {
  "main": {
    "App": {
		EncodeImage(arg1:string):Promise<string>
		GenerateCollectionFromConfig(arg1:models.NewCollectionConfig):Promise<void>
		GenerateNewCollectionFromConfig(arg1:models.NewCollectionConfig):Promise<void>
		GetApplicationDocumentsDirectory():Promise<string>
		GetImageStats(arg1:string):Promise<models.FileInfo>
		GetPreview(arg1:models.NewCollectionConfig):Promise<string>
		GetSettings():Promise<models.Settings>
		MessageDialog(arg1:models.MessageDialogOptions):Promise<string>
		OpenDirectoryDialog():Promise<string>
		OpenFileDialog():Promise<string>
		ReadLayers(arg1:string):Promise<models.CollectionConfig>
		SaveDocuments(arg1:Array<models.Document>):Promise<void>
		SaveFile(arg1:string,arg2:string):Promise<boolean>
		SaveFileDialog():Promise<string>
		SaveSettings(arg1:models.Settings):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
