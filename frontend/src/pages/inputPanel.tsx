// React Component
import { Button } from "react-bootstrap";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";

// Project Specified Component
import FolderNameDisplay from "@/components/FolderDisplay";
import { ModalControl } from "@/controls/ModalControl";
import BookMetadata from "@/pages/metadata/BookMetadata";
import CreatorMetadata from "@/pages/metadata/CreatorMetadata";
import MiscMetadata from "@/pages/metadata/MiscMetadata";
import SeriesMetadata from "@/pages/metadata/SeriesMetadata";
import TagMetadata from "@/pages/metadata/TagMetadata";

// Wails binding
import { ExportXml } from "@wailsjs/go/application/App";
import { comicinfo } from "@wailsjs/go/models";

/** Props Interface for InputPanel */
type InputProps = {
    /** The comic info object. Accept undefined value. */
    comicInfo: comicinfo.ComicInfo | undefined;

    /** Complete folder path, for reference. */
    folderPath?: string;

    /** The function to change display panel to export panel. */
    toExport: () => void;

    /** The info Setter. This function should be doing setting value, but no verification. */
    infoSetter: (field: string, value: string | number) => void;

    /** Modal Controller. */
    modalControl: ModalControl;
};

/**
 * The panel for input/edit content of ComicInfo.xml
 * @returns JSX Element
 */
export default function InputPanel({
    comicInfo,
    folderPath,
    toExport,
    infoSetter,
    modalControl,
}: Readonly<InputProps>) {
    /** Save current comic information to xml file. */
    function save() {
        if (folderPath === undefined) {
            modalControl.showErr("Folder path is not defined. Please try again.");
            return;
        }

        if (comicInfo === undefined) {
            modalControl.showErr("Empty comicinfo. Please try again.");
            return;
        }

        // Start Running
        ExportXml(folderPath, comicInfo).then((msg) => {
            console.log(`xml return: '${msg}'`);
            if (msg === "") {
                modalControl.complete();
            } else {
                modalControl.showErr(msg);
            }
        });
    }

    return (
        <div id="Input-Panel" className="mt-5">
            <h5 className="mb-4">Modify ComicInfo.xml</h5>

            {/* Component for showing folder name (with basename only) */}
            <FolderNameDisplay folderPath={folderPath} />

            {/* The Tabs Group to display metadata. */}
            <Tabs defaultActiveKey="Main" id="uncontrolled-tab-example" className="mb-3">
                <Tab eventKey="Main" title="Book Metadata">
                    <BookMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
                </Tab>

                <Tab eventKey="Creator" title="Creator">
                    <CreatorMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
                </Tab>

                <Tab eventKey="Tags" title="Tags">
                    <TagMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
                </Tab>

                <Tab eventKey="Series" title="Series">
                    <SeriesMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
                </Tab>

                <Tab eventKey="Misc" title="Collection & ReadList">
                    <MiscMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
                </Tab>
            </Tabs>

            {/* The button that will always at the bottom of screen. Should ensure there has enough space */}
            <div className="fixed-bottom mb-3">
                <Button variant="outline-light" className="mx-2" id="btn-save" onClick={save}>
                    Save
                </Button>
                <Button variant="success" className="mx-2 " id="btn-export-cbz" onClick={toExport}>
                    Export to .cbz
                </Button>
            </div>
        </div>
    );
}
