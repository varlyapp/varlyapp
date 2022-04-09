export enum LoggerType {
    Debug = 'Debug',
    Error = 'Error',
    Fatal = 'Fatal',
    Info = 'Info',
    Trace = 'Trace',
    Warning = 'Warning',
}

type Layer = {
    name: string
    path: string
    width: number
    height: number
    weight: number
    type: string
}

type Trait = {
    name: string;
    collapsed: boolean;
}

type Collection = {
    name: string
    description: string
    sourceDirectory: string
    outputDirectory: string
    traits: Trait[] | Object[]
    layers: {[key: string]: Layer[]} | Object
    width: number
    height: number
    size: number
}

export type { Layer, Trait, Collection }
