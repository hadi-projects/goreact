import axios from './client';

const API_PATH = '/api/v1/popo';

export const getAllPopos = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getPopoById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createPopo = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updatePopo = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deletePopo = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
