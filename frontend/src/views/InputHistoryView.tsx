import { type ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";

import DataTable from "@/components/DataTable";
import DataTableRowAction from "@/components/DataTableRowAction";
import { GetAllAutofillWord } from "@wailsjs/go/application/App";
import type { store } from "@wailsjs/go/models";

const columnDef: ColumnDef<store.AutofillWord>[] = [
    { accessorKey: "word", header: "Word" },
    { accessorKey: "category", header: "Category" },
    {
        id: "action",
        cell: ({ row }) => {
            const word = row.original;
            return (
                <DataTableRowAction
                    data={word}
                    editAction={() => console.log(`edit ${word.id}`)}
                    deleteAction={() => console.log(`delete ${word.id}`)}
                />
            );
        },
    },
];

const columnClass: string[] = ["w-65", "w-25", "w-10"];

export default function AutofillView() {
    const [data, setData] = useState<store.AutofillWord[]>([]);

    useEffect(() => {
        GetAllAutofillWord().then((resp) => {
            setData(resp);
        });
    }, []);

    return (
        <div id="Autofill-Panel" className="mt-2 mx-5">
            <h5 className="mb-4">Autofill Setting</h5>

            <DataTable columns={columnDef} headerClassName={columnClass} data={data} />
        </div>
    );
}
