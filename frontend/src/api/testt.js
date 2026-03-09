import axios from './client';

const API_PATH = '/api/v1/testt';

export const getAllTestts = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getTesttById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createTestt = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateTestt = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteTestt = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
