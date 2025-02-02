import { env } from '$env/dynamic/private';

const BACKEND_URL = env.BACKEND_URL;

import axios from "axios";
import noImg from '../../assets/no-img.png';
import { Buffer } from 'buffer';

export default class DatabaseFetcher {
    async getAllPhotos() {
        try {
            const response = await axios.get(`${BACKEND_URL}/files`);
            const photos = response.data;
            
            const imgPromises = photos
                .map(async (photo: { url: string; }) => {
                    const photoId = photo.url.split('/').pop();
                    if (photoId && photoId.includes('example')) {
                        return noImg;
                    } else {
                        const imgResponse = await axios.get(
                            `${BACKEND_URL}/files/${photoId}`,
                            { responseType: 'arraybuffer' } 
                        );
                        
                        const base64 = Buffer.from(imgResponse.data).toString('base64');
                        const contentType = imgResponse.headers['content-type'];
                        return `data:${contentType};base64,${base64}`;
                    }
                });

            const imgs = await Promise.all(imgPromises);

            return {
                photos: photos,
                imgs: imgs 
            };
        } catch (error) {
            console.error('Error fetching photos:', error);
            throw error;
        }
        
    }

    async uploadPhoto(photo: Blob, title: string, description: string) {
        const formData = new FormData();
        formData.append('image', photo);
        formData.append('title', title);
        formData.append('description', description);

        try {
            const response = await axios.post(`${BACKEND_URL}/upload`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });

            if (response.statusText == "OK") {
                console.log("Successfully uploaded!");
                return { success: true };
            } else {
                console.log("Failed to upload" + await response.data);
                return { success: false, message: 'Upload failed.' + (await response.data) };
            }
        } catch (error) {
            console.error("Error during upload:", error);
            return { success: false, message: 'Upload failed.' };
        }
    }

    async deletePhoto(photoId: string) {
        try {
            const response = await axios.delete(`${BACKEND_URL}/files/${photoId}`);
            if (response.statusText == "OK") {
                console.log("Successfully deleted!");
                return { success: true };
            } else {
                console.log("Failed to delete" + await response.data);
                return { success: false, message: 'Delete failed.' + (await response.data) };
            }
        } catch (error) {
            console.error("Error during delete:", error);
            return { success: false, message: 'Delete failed.' };
        }
    }
}
