import { PencilLine, Trash2 } from "lucide-react";

type DataTableRowActionProps<DataT> = {
    data: DataT;
    editAction: () => void;
    deleteAction: () => void;
};

export default function DataTableRowAction<DataT>(props: Readonly<DataTableRowActionProps<DataT>>) {
    return (
        <>
            <PencilLine className="mx-1" onClick={props.editAction} size={20} role="button" />
            <Trash2 className="mx-1" onClick={props.deleteAction} size={20} color="red" role="button" />
        </>
    );
}
