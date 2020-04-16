import React from "react";
import { Card } from "react-bootstrap";
import { toast } from "react-toastify";

export const noop = () => {};

export const Cards = (heading, text) => {
  return (
    <Card bg="info" text="white">
      <Card.Header>{heading}</Card.Header>
      <Card.Body>
        <Card.Text>{text}</Card.Text>
      </Card.Body>
    </Card>
  );
};

export const showToastr = (type, ...rest) => {
  toast[type](...rest);
};

export const successErrorHandler = (resolve, reject) => {
  const success = (data, status) => resolve(data);
  const err = error => {
    reject && reject(error);
  };
  return {
    success,
    err
  };
};

export const showToastrError = errObj => {
  showToastr("error", errObj.error || errObj.err || "Something went wrong.", null, {
    timeOut: 0,
    extendedTimeOut: 0
  });
};