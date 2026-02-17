import axios from './client';

const API_PATH = '/api/v1/sdsdsdsdd';

export const getAllSdsdsds = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getSdsdsdById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createSdsdsd = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateSdsdsd = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteSdsdsd = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
