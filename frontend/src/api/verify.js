import axios from './client';

const API_PATH = '/api/v1/verifies';

export const getAllVerifys = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getVerifyById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createVerify = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateVerify = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteVerify = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
