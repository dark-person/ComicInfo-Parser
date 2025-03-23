// React Component
import { useState } from "react";
import { Col, Row } from "react-bootstrap";
import Button from "react-bootstrap/Button";

// Project Specified Component
import { AppMode } from "@/controls/AppMode";
import { ModalControl } from "@/controls/ModalControl";
import { ExportMethod, SessionData } from "@/controls/SessionData";
import ExportPanel from "@/pages/exportPanel";
import FolderSelect from "@/pages/folderSelect";
import HelpPanel from "@/pages/helpPanel";
import InputPanel from "@/pages/inputPanel";

// Wails
import { GetComicInfo } from "@wailsjs/go/application/App";
import { comicinfo } from "@wailsjs/go/models";

type CreateCbzViewProps = {
    /** Decide which panel will be displayed. */
    mode: AppMode;
    /** Hooks for decide which panel will be displayed. */
    setMode: (mode: AppMode) => void;
    /** Modal State to control display which dialog. */
    modalController: ModalControl;
};

/** Panel to create a cbz file. */
export default function CreateCbzView({ mode, setMode, modalController }: Readonly<CreateCbzViewProps>) {
    /** The ComicInfo model. For communicate with different panel. */
    const [info, setInfo] = useState<comicinfo.ComicInfo | undefined>(undefined);

    /** The directory of initial input, which is the folder contain image. */
    const [inputDir, setInputDir] = useState<string | undefined>(undefined);

    const [sessionData, setSessionData] = useState<SessionData>({
        exportMethod: ExportMethod.DEFAULT_WRAP_CBZ,
        deleteAfterExport: false,
    });

    /**
     * Set value of selected folder, then pass selected folder to next panel.
     * @param folder the absolute path to the folder
     */
    function passingFolder(folder: string) {
        console.log("passing folder: " + folder);

        // Set Loading Modal
        modalController.loading();

        // Get ComicInfo
        GetComicInfo(folder).then((response) => {
            const error = response.ErrorMessage;

            if (error !== "") {
                // Print Error Message
                modalController.showErr(error);
            } else {
                // Reset all modal
                modalController.closeAll();

                // Set data with info
                setInfo(response.ComicInfo);
                setInputDir(folder);

                // Pass to another panel
                setMode(AppMode.INPUT_DATA);
            }
        });
    }

    /** Change the panel in app to export panel. */
    function showExportPanel() {
        setMode(AppMode.EXPORT);
    }

    /** Show help panel in App. */
    function showHelpPanel() {
        setMode(AppMode.HELP);
    }

    /**
     * Return to previous page.
     * Only `AppMode.INPUT_DATA` & `AppMode.EXPORT` is supported.
     */
    function backward() {
        switch (mode) {
            case AppMode.INPUT_DATA:
                setMode(AppMode.SELECT_FOLDER);
                return;
            case AppMode.EXPORT:
                setMode(AppMode.INPUT_DATA);
                return;
            default:
                throw new Error("Invalid mode");
        }
    }

    /** Return to the home panel. In current version, it is select folder panel. */
    function backToHomePanel() {
        setMode(AppMode.SELECT_FOLDER);
    }

    /**
     * Set the value by its field name.
     * @param data the data to be modify
     * @param key  the field name
     * @param value new value of that field
     */
    function setValue<T, K extends keyof T>(data: T, key: K, value: T[K]) {
        data[key] = value;
    }

    /**
     * Setter for changing value of comicInfo.
     *
     * Note that this function will treat "Summary" field as special case.
     * @param field the field name, must be same as ComicInfo field name
     * @param value new value of that field
     */
    function infoSetter(field: string, value: string | number) {
        // Prepare an object of ComicInfo
        const temp = { ...info } as comicinfo.ComicInfo;

        // Treat Summary field name as special case
        if (field === "Summary" && typeof value === "string") {
            temp["Summary"]["InnerXML"] = value;
        } else {
            // Normal Change
            const key = field as keyof comicinfo.ComicInfo;
            setValue(temp, key, value);
        }

        // Set the changed value to data
        setInfo(temp);
    }

    return (
        <Row className="min-vh-100">
            {/* Back Button, return to previous panel */}
            <Col xs={1} className="mt-4">
                {/* Only Allow backward when export page / input data page */}
                {(mode === AppMode.EXPORT || mode === AppMode.INPUT_DATA) && (
                    <Button variant="secondary" onClick={backward}>
                        {"<"}
                    </Button>
                )}
            </Col>

            {/* Area to display panel */}
            <Col>
                {mode === AppMode.SELECT_FOLDER && (
                    <FolderSelect
                        handleFolder={passingFolder}
                        showHelpPanel={showHelpPanel}
                        modalControl={modalController}
                    />
                )}
                {mode === AppMode.INPUT_DATA && (
                    <InputPanel
                        comicInfo={info}
                        toExport={showExportPanel}
                        infoSetter={infoSetter}
                        folderPath={inputDir}
                        modalControl={modalController}
                    />
                )}
                {mode === AppMode.EXPORT && (
                    <ExportPanel
                        comicInfo={info}
                        originalDirectory={inputDir}
                        modalControl={modalController}
                        exportMethod={sessionData.exportMethod}
                        setExportMethod={(val) => setSessionData({ ...sessionData, exportMethod: val })}
                        deleteAfterExport={sessionData.deleteAfterExport}
                        setDeleteAfterExport={(val) => setSessionData({ ...sessionData, deleteAfterExport: val })}
                    />
                )}
                {mode === AppMode.HELP && <HelpPanel backToHome={backToHomePanel} />}
            </Col>

            {/* Use as alignment */}
            <Col xs={1} className="align-self-center"></Col>
        </Row>
    );
}
