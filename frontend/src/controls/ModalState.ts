/** State to be used to control modal located at App.tsx. */
export type ModalState = {
    /** True if need to display loading dialog. Should be show when change to another page. */
    isLoading: boolean;
    /** Error Message, will modal be display when not empty string. Empty Strings mean not error at all. */
    errMsg: string;
    /** True if completed screen should appear */
    isCompleted: boolean;
    /** True if back home when modal dispose, false otherwise. */
    resetOnComplete: boolean;
};

/** Default status of ModalState, which means no modal shown. */
export const defaultModalState: ModalState = {
    isLoading: false,
    errMsg: "",
    isCompleted: false,
    resetOnComplete: false,
};
