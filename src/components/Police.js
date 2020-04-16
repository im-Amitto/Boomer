import React from "react";
import DataTable from "react-data-table-component";
import { Button, Form, InputGroup } from "react-bootstrap";
import { getAllOfficer, addNewOfficer } from "../common/actions";
import { showToastr } from "../common/utils";

class Police extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      filterText: "",
      resetPaginationToggle: false,
      data: [],
      filteredData: [],
    };
    this.fetchData();
  }

  fetchData() {
    getAllOfficer().then((cases) => {
      this.setState({
        data: cases,
        filteredData: cases,
        filterText: "",
      });
    });
  }

  handleChange = (e) => {
    this.filterItems(e.target.value);
  };

  handleClear = () => {
    if (this.state.filterText) {
      this.setState(
        {
          filterText: "",
          resetPaginationToggle: !this.state.resetPaginationToggle,
        },
        () => {
          this.filterItems(this.state.filterText);
        }
      );
    }
  };

  filterItems = (filterText) => {
    let filteredData = this.state.data;
    if (filterText !== "") {
      filteredData = filteredData.filter(
        (item) =>
          item.name &&
          item.name.toLowerCase().includes(filterText.toLowerCase())
      );
    }
    this.setState({
      filterText,
      filteredData,
    });
  };

  columns = [
    {
      name: "Name",
      selector: "name",
      sortable: true,
    },
    {
      name: "Last Active",
      selector: "lastactive",
      sortable: true,
      right: true,
    },
    {
      name: "working",
      selector: "image",
      cell: (d) => (d.status ? "Yes" : "No"),
    },
  ];

  handleSubmit = (event) => {
    event.preventDefault();
    let payload = { name: event.target.elements.officerName.value };
    addNewOfficer(payload).then((data) => {
      showToastr("success", data["msg"]);
      this.fetchData();
    });
  };

  render() {
    return (
      <div className="container">
        <div className="row mt-5">
          <div className="cl-sm-4  text-center offset-sm-4">
            <Form onSubmit={(e) => this.handleSubmit(e)}>
              <Form.Group>
                <Form.Control
                  name="officerName"
                  type="text"
                  placeholder="Enter officer name"
                />
                <Form.Text className="text-muted">
                  We'll never share your name with anyone else.
                </Form.Text>
              </Form.Group>
              <Button variant="primary" type="submit">
                Add new officer
              </Button>
            </Form>
          </div>
        </div>
        <div className="row mt-3">
          <div className="col-sm-4 offset-sm-8">
            <Form.Group>
              <InputGroup>
                <Form.Control
                  type="text"
                  placeholder="Filter by Name"
                  aria-describedby="inputGroupPrepend"
                  name="username"
                  value={this.state.filterText}
                  onChange={this.handleChange}
                />
                <InputGroup.Append>
                  <Button variant="danger" onClick={this.handleClear}>
                    X
                  </Button>
                </InputGroup.Append>
              </InputGroup>
            </Form.Group>
          </div>
        </div>
        <DataTable
          title="Manage Cases"
          columns={this.columns}
          data={this.state.filteredData}
          highlightOnHover
          defaultSortField="name"
          pagination
          paginationResetDefaultPage={this.state.resetPaginationToggle}
          persistTableHead
        />
      </div>
    );
  }
}

export default Police;
