// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import OptionFormRow from "@/forms/OptionFormRow";
import { type MetadataProps } from "@/pages/metadata/MetadataProps";
import { HistoryResp } from "@bindings/services";

/**
 * Dummy function to return empty options for default options in `OptionFormRow`.
 * @returns promise to return an empty `HistoryResp`
 */
function emptyOptions(): Promise<HistoryResp> {
    return new Promise<HistoryResp>((resolve) => {
        const resp = new HistoryResp();
        resp.Inputs = [];
        resp.ErrorMsg = "";

        resolve(resp);
    });
}

/** The JSX Element for show/edit entity metadata. */
export default function EntityMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
    return (
        <div>
            <Form>
                <OptionFormRow
                    title={"Characters"}
                    value={info?.Characters}
                    setValue={(val) => infoSetter("Characters", val)}
                    getDefaultOpt={emptyOptions}
                    componentHeight="120px"
                />

                <OptionFormRow
                    title={"Teams"}
                    value={info?.Teams}
                    setValue={(val) => infoSetter("Teams", val)}
                    getDefaultOpt={emptyOptions}
                    componentHeight="120px"
                />

                <OptionFormRow
                    title={"Locations"}
                    value={info?.Locations}
                    setValue={(val) => infoSetter("Locations", val)}
                    getDefaultOpt={emptyOptions}
                    componentHeight="120px"
                />
            </Form>
        </div>
    );
}
