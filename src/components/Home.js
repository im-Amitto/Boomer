import React from "react";
import { Button, Form } from "react-bootstrap";
import { Cards, showToastr } from "../common/utils";
import { getAllCases, getAllOfficer, addNewCase } from "../common/actions";

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      totalCases: 0,
      totalOfficers: 0,
      casesSolved: 0,
    };
    this.getData();
  }
  getData() {
    getAllCases().then((cases) => {
      getAllOfficer().then((officers) => {
        let solvedCases = 0;
        cases.forEach((o) => {
          if (o.status) {
            solvedCases++;
          }
        });
        this.setState({
          totalCases: cases.length,
          totalOfficers: officers.length,
          casesSolved: solvedCases,
        });
      });
    });
  }

  handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData();
    data.append("myFile", event.target.file.files[0]);
    data.append("vehicleNumber", event.target.elements.vehicleNumber.value);
    addNewCase(data).then(data=>{
      showToastr("success", data['msg']);
      this.getData();
    })
  };

  render() {
    return (
      <div className="container-fluid text-center">
        <h2 className="mt-5">Welcome to Boomer</h2>
        <h5>We find your vehicle ASAP</h5>
        <div className="row mt-5">
          <div className="col-sm-4 py-1">
            {Cards("Total cases", this.state.totalCases)}
          </div>
          <div className="col-sm-4 py-1">
            {Cards("Total Officers", this.state.totalOfficers)}
          </div>
          <div className="col-sm-4 py-1">
            {Cards("Cases Solved", this.state.casesSolved)}
          </div>
        </div>
        <div className="row mt-5">
          <div className="col-sm-4 offset-sm-4">
            <Form onSubmit={(e) => this.handleSubmit(e)}>
              <Form.Group>
                <Form.Control
                  type="text"
                  placeholder="Enter vehicle number"
                  name="vehicleNumber"
                  required
                />
                <Form.Text className="text-muted">
                  We'll never share your vehicle number with anyone else.
                </Form.Text>
              </Form.Group>
              <Form.Group>
                <Form.File
                  name="file"
                  id="custom-file"
                  label="Vehicle Image"
                  data-browse="Upload"
                  custom
                  required
                />
              </Form.Group>
              <Button variant="warning" type="submit">
                Report vehicle
              </Button>
            </Form>
          </div>
        </div>
      </div>
    );
  }
}

export default Home;
