// CSS Import
import "@/App.css";
import "bootstrap/dist/css/bootstrap.min.css";

// React Component
import { useState } from "react";
import { Tab, Tabs } from "react-bootstrap";

// Project Specified Component
import { CompleteModal, ErrorModal, LoadingModal } from "@/components/modal";
import { AppMode } from "@/controls/AppMode";
import { ModalControl } from "@/controls/ModalControl";
import { defaultModalState, ModalState } from "@/controls/ModalState";
import CreateCbzView from "@/views/CreateCbzView";

/**
 * The main component to be displayed. It will handle all pages data & timing to display.
 *
 * There will have two column with width=1 at left and right side, for center the panel/page component.
 *
 * There has a return button at the top-left corner in this window.
 */
function App() {
    const [createMode, setCreateMode] = useState<AppMode>(AppMode.SELECT_FOLDER);

    /** Modal State to control display which dialog. */
    const [modalState, setModalState] = useState<ModalState>(defaultModalState);

    /** Controller of modal. */
    const modalController: ModalControl = {
        showErr: (err) => setModalState({ ...defaultModalState, errMsg: err }),
        loading: () => setModalState({ ...defaultModalState, isLoading: true }),
        complete: () => setModalState({ ...defaultModalState, isCompleted: true }),
        completeAndReset: () => setModalState({ ...defaultModalState, isCompleted: true, resetOnComplete: true }),
        closeAll: () => setModalState({ ...defaultModalState }),
    };

    /** Return to the home. In current version, it is select folder panel. */
    function backToCreateHome() {
        setCreateMode(AppMode.SELECT_FOLDER);
    }

    return (
        <div id="App" className="container-fluid">
            {/* Modal Part */}
            <LoadingModal show={modalState.isLoading} />

            <ErrorModal
                show={modalState.errMsg !== ""}
                errorMessage={modalState.errMsg}
                disposeFunc={modalController.closeAll}
            />

            <CompleteModal
                show={modalState.isCompleted}
                disposeFunc={() => {
                    modalController.closeAll();

                    // Redirect to first page only if export cbz is clicked
                    if (modalState.resetOnComplete) {
                        backToCreateHome();
                    }
                }}
            />

            {/* Main Panel of this app */}
            <Tabs className="mt-2 mb-3">
                <Tab eventKey={"create"} title={"Create CBZ"}>
                    <CreateCbzView mode={createMode} setMode={setCreateMode} modalController={modalController} />
                </Tab>
            </Tabs>
        </div>
    );
}

export default App;
