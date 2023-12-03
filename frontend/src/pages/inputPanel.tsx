// React Component
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";

export default function InputPanel() {
	return (
		<div id="Input-Panel" className="mt-5">
			<h5 className="mb-4">Modify ComicInfo.xml</h5>
			<Tabs
				defaultActiveKey="Main"
				id="uncontrolled-tab-example"
				className="mb-3">
				<Tab eventKey="Main" title="Main">
					Tab content for Home
				</Tab>
				<Tab eventKey="Creator" title="Creator">
					Tab content for Profile
				</Tab>
				<Tab eventKey="Tags" title="Tags" disabled>
					Tab content for Contact
				</Tab>
			</Tabs>
		</div>
	);
}
