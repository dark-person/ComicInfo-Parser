import { Button } from "react-bootstrap";
import "./Tags.css";

/** The props of `TagsArea` */
type TagsAreaProps = {
    /** Raw string of tags, which is only separate tag with comma. */
    rawTags: string | undefined;
    /** The function to delete tag from `comicInfo`. */
    handleDelete: (arg0: number) => void;
};

/** Area that holding for Tags. */
export function TagsArea({ rawTags, handleDelete }: Readonly<TagsAreaProps>) {
    /**
     * Parse the raw string that contains tags into arrays.
     * <p>
     * Please note that, split may able to return `[""]` strings array,
     * which may result generated a empty tag item.
     * Therefore, this function will turn this case in to empty array instead.
     *
     * @param raw raw string of tags, e.g `"tag1, tag2"`
     * @returns array of tags, e.g. `["tag1", "tag2"]`
     */
    function getTagsList(raw: string | undefined): string[] {
        if (raw === undefined) {
            return [];
        }

        const temp = raw.split(",");

        // Prevent Empty Tag list
        if (temp.length === 1 && temp[0].length === 0) {
            console.log("empty tag list");
            return [];
        }

        return temp;
    }

    return (
        <div className="tag-area text-start p-2  ">
            <div className="d-inline-flex flex-wrap">
                {getTagsList(rawTags).map((item, index) => (
                    <Tag tag={item} key={"tags-" + index} index={index} handleDelete={handleDelete} />
                ))}
            </div>
        </div>
    );
}

/** The props of `Tag` component. */
type TagProps = {
    /** The value of tag name. */
    tag: string;
    /** The index of tag, also represent index of tags array. */
    index: number;
    /** The function to delete tag from `comicInfo`. */
    handleDelete: (arg0: number) => void;
};

/** Item for holding one Tag, include delete button in every tag. */
export function Tag({ tag, index, handleDelete }: Readonly<TagProps>) {
    return (
        <div className="me-1 mb-1 tag-item bg-secondary p-1">
            <div className="d-flex align-items-center">
                <span className="px-1">{tag}</span>
                <Button className="btn-close remove-icon" onClick={() => handleDelete(index)}></Button>
            </div>
        </div>
    );
}
