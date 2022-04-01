// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "main": {
    "App": {
      /**
       * EncodeImage
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "EncodeImage": (arg1) => {
        return window.go.main.App.EncodeImage(arg1);
      },
      /**
       * GenerateCollectionFromConfig
       * @param {models.NewCollectionConfig} arg1 - Go Type: nft.NewCollectionConfig
       * @returns {Promise<void>} 
       */
      "GenerateCollectionFromConfig": (arg1) => {
        return window.go.main.App.GenerateCollectionFromConfig(arg1);
      },
      /**
       * GenerateNewCollectionFromConfig
       * @param {models.NewCollectionConfig} arg1 - Go Type: nft.NewCollectionConfig
       * @returns {Promise<void>} 
       */
      "GenerateNewCollectionFromConfig": (arg1) => {
        return window.go.main.App.GenerateNewCollectionFromConfig(arg1);
      },
      /**
       * GetApplicationDocumentsDirectory
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetApplicationDocumentsDirectory": () => {
        return window.go.main.App.GetApplicationDocumentsDirectory();
      },
      /**
       * GetImageStats
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<models.FileInfo>}  - Go Type: fs.FileInfo
       */
      "GetImageStats": (arg1) => {
        return window.go.main.App.GetImageStats(arg1);
      },
      /**
       * GetSettings
       * @returns {Promise<models.Settings>}  - Go Type: *settings.Settings
       */
      "GetSettings": () => {
        return window.go.main.App.GetSettings();
      },
      /**
       * MessageDialog
       * @param {models.MessageDialogOptions} arg1 - Go Type: frontend.MessageDialogOptions
       * @returns {Promise<string>}  - Go Type: string
       */
      "MessageDialog": (arg1) => {
        return window.go.main.App.MessageDialog(arg1);
      },
      /**
       * OpenDirectoryDialog
       * @returns {Promise<string>}  - Go Type: string
       */
      "OpenDirectoryDialog": () => {
        return window.go.main.App.OpenDirectoryDialog();
      },
      /**
       * OpenFileDialog
       * @returns {Promise<string>}  - Go Type: string
       */
      "OpenFileDialog": () => {
        return window.go.main.App.OpenFileDialog();
      },
      /**
       * ReadLayers
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<models.CollectionConfig>}  - Go Type: fs.CollectionConfig
       */
      "ReadLayers": (arg1) => {
        return window.go.main.App.ReadLayers(arg1);
      },
      /**
       * SaveDocuments
       * @param {Array<models.Document>} arg1 - Go Type: []settings.Document
       * @returns {Promise<void>} 
       */
      "SaveDocuments": (arg1) => {
        return window.go.main.App.SaveDocuments(arg1);
      },
      /**
       * SaveFile
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @returns {Promise<boolean>}  - Go Type: bool
       */
      "SaveFile": (arg1, arg2) => {
        return window.go.main.App.SaveFile(arg1, arg2);
      },
      /**
       * SaveFileDialog
       * @returns {Promise<string>}  - Go Type: string
       */
      "SaveFileDialog": () => {
        return window.go.main.App.SaveFileDialog();
      },
      /**
       * SaveSettings
       * @param {models.Settings} arg1 - Go Type: *settings.Settings
       * @returns {Promise<void>} 
       */
      "SaveSettings": (arg1) => {
        return window.go.main.App.SaveSettings(arg1);
      },
    },
  },

};
export default go;
