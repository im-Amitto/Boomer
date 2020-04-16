import { BASE_URL } from "./constatns";
import { getRequest, postRequest, performMultipartRequest } from "../api";

export const casesUrl =  `${BASE_URL}/case`;
export const officerUrl =  `${BASE_URL}/officer`;
export const caseCompleteUrl  =  `${BASE_URL}/caseComplete`;

export const addNewCase = payload => {
  const url = casesUrl;
  return performMultipartRequest(url, "post", payload);
};
export const addNewOfficer = payload => {
  const url = officerUrl;
  return postRequest(url, payload);
};
export const caseCompleted = payload => {
  const url = caseCompleteUrl;
  return postRequest(url, payload);
};
export const getAllOfficer = () => {
  const url = officerUrl;
  return getRequest(url);
};
export const getAllCases = () => {
  const url = casesUrl;
  return getRequest(url);
};