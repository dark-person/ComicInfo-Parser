import { ColoredCheckBox } from "@/components/ColoredFormCheck";
import FolderSelector from "@/components/FolderSelector";
import {
    AutoGenerateXML,
    GetComicFolder,
} from "@bindings/github.com/dark-person/comicinfo-parser/internal/services/comicservice";
import { useEffect, useState } from "react";
import { Button } from "react-bootstrap";

type AutomationState = {
    /** The Directory Absolute Path selected by User. */
    directory: string;
    isOverwrite: boolean;
};

/** View for automation jobs. */
export default function AutomationView() {
    /** Default directory for comic folder. */
    const [comicFolder, setComicFolder] = useState<string | undefined>(undefined);

    const [state, setState] = useState<AutomationState>({
        directory: "",
        isOverwrite: false,
    });

    useEffect(() => {
        void GetComicFolder().then((dir: string) => setComicFolder(dir));
    }, []);

    function startAutomation() {
        AutoGenerateXML(state.directory, state.isOverwrite)
            .then((resp) => {
                console.log(JSON.stringify(resp, null, 4));
            })
            .catch((err) => console.error(err));
    }

    return (
        <div id="Automation-Panel" className="mt-2 mx-5">
            <h5 className="mb-4">Automation</h5>

            <p>This will help you to auto generate ComicInfo.xml as possible.</p>

            {/* Folder Chooser */}
            <FolderSelector
                id="automation-folder-select"
                className="mb-3 w-75 mx-auto"
                label="Input Folder"
                buttonId="btn-select-automation-folder"
                directory={state.directory}
                setDirectory={(str) => setState({ ...state, directory: str })}
                defaultDirectory={comicFolder}
            />

            {/* Options */}
            <div className="w-25 mx-auto d-grid gap-2 mt-3">
                <ColoredCheckBox
                    id="overwrite-exist-xml"
                    name="overwrite-exist-xml"
                    color="dark-red"
                    label="Overwrite existing XML"
                    checked={state.isOverwrite}
                    onChange={() => setState({ ...state, isOverwrite: !state.isOverwrite })}
                />
            </div>

            {/* Button Group */}
            <div className="w-25 mx-auto d-grid gap-2">
                <Button variant="outline-info" id="btn-quick-export" onClick={startAutomation}>
                    Start Automation
                </Button>
            </div>
        </div>
    );
}
