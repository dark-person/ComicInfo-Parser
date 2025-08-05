import DataTable from "@/components/DataTable";
import { type ColumnDef } from "@tanstack/react-table";

export type AutofillWord = {
    id: number;
    word: string;
    category: string;
};

const ExampleData: AutofillWord[] = Array.from<AutofillWord>({ length: 50 });

for (let i = 0; i < 25; i++) {
    ExampleData[i] = { id: i + 1, word: `word${i + 1}`, category: "tag" };
    ExampleData[i + 25] = { id: i + 26, word: `word${i + 26}`, category: "genre" };
}

const columnDef: ColumnDef<AutofillWord>[] = [
    { accessorKey: "word", header: "Word" },
    { accessorKey: "category", header: "Category" },
];

export default function AutofillView() {
    return (
        <div id="Autofill-Panel" className="mt-2 mx-5">
            <h5 className="mb-4">Autofill Setting</h5>

            <DataTable columns={columnDef} data={ExampleData} />
        </div>
    );
}
