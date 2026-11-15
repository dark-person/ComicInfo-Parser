import type { ComicInfo } from "@bindings/comicinfo/models";

/** The Props for Metadata Component. */
export type MetadataProps = {
    /** The comic info object. Accept undefined value. */
    comicInfo: ComicInfo | undefined;

    /** The info Setter. This function should be doing setting value, but no verification. */
    infoSetter: (field: string, value: string | number) => void;
};
