// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import OptionFormRow from "@/forms/OptionFormRow";
import { type MetadataProps } from "@/pages/metadata/MetadataProps";
import { GetAllTagInput } from "@bindings/github.com/dark-person/comicinfo-parser/internal/services/comicservice";

// Wails binding

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
export default function TagMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
    return (
        <div>
            <Form>
                <OptionFormRow
                    title={"Tags"}
                    value={info?.Tags}
                    setValue={(val) => infoSetter("Tags", val)}
                    getDefaultOpt={GetAllTagInput}
                    componentHeight="200px"
                    menuMaxHeight="16em"
                />
            </Form>
        </div>
    );
}
