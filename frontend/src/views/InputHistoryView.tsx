import { type ColumnDef } from "@tanstack/react-table";
import { useEffect, useState } from "react";

import DataTable from "@/components/DataTable";
import { GetAllAutofillWord } from "@bindings/github.com/dark-person/comicinfo-parser/internal/services/comicservice";
import type { AutofillWord } from "@bindings/github.com/dark-person/comicinfo-parser/internal/store/models";

const columnDef: ColumnDef<AutofillWord>[] = [
    { accessorKey: "word", header: "Word" },
    { accessorKey: "category", header: "Category" },
];

const columnClass: string[] = ["w-65", "w-35"];

export default function AutofillView() {
    const [data, setData] = useState<AutofillWord[]>([]);

    useEffect(() => {
        GetAllAutofillWord().then((resp: AutofillWord[]) => {
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
