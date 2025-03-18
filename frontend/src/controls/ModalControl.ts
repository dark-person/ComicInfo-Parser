/** A interface for accessing modal message in App component. */
export interface ModalControl {
    /**
     * Close other modal, and show error message modal when `message` is not empty.
     * Otherwise, remove error message modal is any.
     *
     * @param message error message to display, or empty string to hide error modal.
     */
    showErr: (message: string) => void;

    /**	  Close other modal, and show loading modal.	 */
    loading: () => void;

    /**	  Close other modal, and show complete modal  */
    complete: () => void;

    /**
     * Close other modal, and show complete modal.
     *
     * After user dispose complete modal, then back to home page.
     * If developer want to remain on current page, then use {@link complete} method instead.
     */
    completeAndReset: () => void;

    /** Close all modal, i.e. reset modal state to default. */
    closeAll: () => void;
}
