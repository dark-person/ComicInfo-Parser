// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import OptionFormRow from "@/forms/OptionFormRow";
import { type MetadataProps } from "@/pages/metadata/MetadataProps";

// Wails binding
import { GetAllCharactersInput, GetAllLocationsInput, GetAllTeamsInput } from "@bindings/services/comicservice";

/** The JSX Element for show/edit entity metadata. */
export default function EntityMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
    return (
        <div>
            <Form>
                <OptionFormRow
                    title={"Characters"}
                    value={info?.Characters}
                    setValue={(val) => infoSetter("Characters", val)}
                    getDefaultOpt={GetAllCharactersInput}
                    componentHeight="120px"
                />

                <OptionFormRow
                    title={"Teams"}
                    value={info?.Teams}
                    setValue={(val) => infoSetter("Teams", val)}
                    getDefaultOpt={GetAllTeamsInput}
                    componentHeight="120px"
                />

                <OptionFormRow
                    title={"Locations"}
                    value={info?.Locations}
                    setValue={(val) => infoSetter("Locations", val)}
                    getDefaultOpt={GetAllLocationsInput}
                    componentHeight="120px"
                />
            </Form>
        </div>
    );
}
