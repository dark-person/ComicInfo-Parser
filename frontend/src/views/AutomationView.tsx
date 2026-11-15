import { useEffect, useState } from "react";
import { Button } from "react-bootstrap";

import FolderSelector from "@/components/FolderSelector";
import { type ModalControl } from "@/controls/ModalControl";

import { AutoGenerateXML, GetComicFolder } from "@bindings/services/comicservice";

type AutomationState = {
    /** The Directory Absolute Path selected by User. */
    directory: string;
};

type AutomationViewProps = {
    /** Modal State to control display which dialog. */
    modalController: ModalControl;
};

/** View for automation jobs. */
export default function AutomationView({ modalController }: Readonly<AutomationViewProps>) {
    /** Default directory for comic folder. */
    const [comicFolder, setComicFolder] = useState<string | undefined>(undefined);

    const [state, setState] = useState<AutomationState>({
        directory: "",
    });

    useEffect(() => {
        void GetComicFolder().then((dir: string) => setComicFolder(dir));
    }, []);

    function startAutoGenerateComicinfo() {
        modalController.loading();

        AutoGenerateXML(state.directory)
            .then((resp) => {
                console.log(JSON.stringify(resp, null, 4));

                // Check if any error occur
                if (resp.errorMessage !== "") {
                    modalController.showErr(resp.errorMessage);
                    return;
                }

                // Otherwise, show complete dialog
                modalController.complete();
            })
            .catch((err) => {
                modalController.showErr(err);
            });
    }

    return (
        <div id="Automation-Panel" className="mt-2 mx-5">
            <h5 className="mb-4">Automation</h5>

            <p>
                This will help you to auto generate ComicInfo.xml as possible. This will only add value if file is
                already exist.
            </p>

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

            {/* Button Group */}
            <div className="w-25 mx-auto d-grid gap-2">
                <Button variant="outline-info" id="btn-quick-export" onClick={startAutoGenerateComicinfo}>
                    Generate Comicinfo.xml
                </Button>
            </div>
        </div>
    );
}
