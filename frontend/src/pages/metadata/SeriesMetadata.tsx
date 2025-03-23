// React
import { ChangeEvent } from "react";

// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import { AgeRatingSelect, MangaSelect } from "@/components/EnumSelect";
import FormRow from "@/forms/FormRow";
import FormSelectRow from "@/forms/FormSelectRow";
import OptionFormRow from "@/forms/OptionFormRow";
import { MetadataProps } from "@/pages/metadata/MetadataProps";

// Wails binding
import { GetAllGenreInput } from "@wailsjs/go/application/App";

export default /** The user interface for show/edit Series MetaData. */
function SeriesMetadata({ comicInfo, infoSetter }: Readonly<MetadataProps>) {
    /**
     * Handler for all input field in this panel.
     * This method will use <code>infoSetter</code> as core,
     * and apply change to comicInfo content.
     * <p>
     * This method will try to find the field is number input first,
     * if field name is number/related, then it will call Number() method before set to ComicInfo
     *
     * @param e the event object, for identify target element
     * @returns void
     */
    function handleChanges(e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>) {
        console.log(e.target.title, e.target.value, e.target.type);

        // Identify Number
        if (e.target.type === "number") {
            infoSetter(e.target.title, Number(e.target.value));
            return;
        }

        // Normal Cases
        infoSetter(e.target.title, e.target.value);
    }

    return (
        <div>
            <Form>
                <FormRow title={"Series"} value={comicInfo?.Series} onChange={handleChanges} />
                <FormRow title={"Volume"} value={comicInfo?.Volume} onChange={handleChanges} />
                <FormRow title={"Count"} value={comicInfo?.Count} onChange={handleChanges} />
                <FormSelectRow
                    title={"AgeRating"}
                    selectElement={<AgeRatingSelect value={comicInfo?.AgeRating} onChange={handleChanges} />}
                />
                <FormSelectRow
                    title={"Manga"}
                    selectElement={<MangaSelect value={comicInfo?.Manga} onChange={handleChanges} />}
                />
                <OptionFormRow
                    title={"Genre"}
                    value={comicInfo?.Genre}
                    setValue={(val) => infoSetter("Genre", val)}
                    getDefaultOpt={GetAllGenreInput}
                />
                <FormRow title={"LanguageISO"} value={comicInfo?.LanguageISO} onChange={handleChanges} />
            </Form>
        </div>
    );
}
