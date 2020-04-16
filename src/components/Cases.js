import React from "react";
import DataTable from "react-data-table-component";
import styled from "styled-components";
import { Button, Form, InputGroup } from "react-bootstrap";
import { getAllCases, caseCompleted } from "../common/actions";
import { showToastr } from "../common/utils";

const SampleStyle = styled.div`
  padding: 16px;
  display: block;
  width: 100%;

  p {
    font-size: 16px;
    font-weight: 700;
    word-break: break-all;
  }
`;

const SampleExpandedComponent = ({ data }) => (
  <SampleStyle>
    <p>{data.summary}</p>
    <img height="75%" width="75%" alt={data.image} src={data.image} />
  </SampleStyle>
);

class Cases extends React.Component {
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
    getAllCases().then((cases) => {
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
          item.vehiclenumber &&
          item.vehiclenumber.toLowerCase().includes(filterText.toLowerCase())
      );
    }
    this.setState({
      filterText,
      filteredData,
    });
  };

  columns = [
    {
      name: "Vehicle Number",
      selector: "vehiclenumber",
      sortable: true,
    },
    {
      name: "Assigned Officer",
      selector: "assignedofficer",
      sortable: true,
      right: true,
    },
    {
      name: "Reported Time",
      selector: "reportedtime",
      sortable: true,
      right: true,
    },
    {
      name: "Image",
      selector: "image",
      cell: (d) => (
        <img height="32x" width="64px" alt={d.image} src={d.image} />
      ),
    },
    {
      name: "Action/Status",
      selector: "image",
      cell: (d) =>
        !d.status ? (
          d.assignedofficer !== 0 ? 
          <Button
            onClick={() => {
              caseCompleted(d).then((o) => {
                showToastr("success",o.msg);
                this.fetchData();
              });
            }}
            variant="success"
            type="submit"
          >
            Vehicle Found
          </Button> : "Waiting fo officer"
        ) : (
          "Case Solved"
        ),
    },
  ];

  render() {
    return (
      <div className="container">
        <div className="row mt-3">
          <div className="col-sm-4 offset-sm-8">
            <Form.Group>
              <InputGroup>
                <Form.Control
                  type="text"
                  placeholder="Filter by Vehicle Number"
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
          expandableRows
          highlightOnHover
          defaultSortField="name"
          expandableRowsComponent={<SampleExpandedComponent />}
          expandOnRowClicked
          pagination
          paginationResetDefaultPage={this.state.resetPaginationToggle}
          persistTableHead
        />
      </div>
    );
  }
}

export default Cases;
