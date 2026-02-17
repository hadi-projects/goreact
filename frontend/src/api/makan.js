import axios from './client';

const API_PATH = '/api/v1/makan';

export const getAllMakans = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getMakanById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createMakan = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateMakan = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteMakan = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
