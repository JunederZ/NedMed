
import axios from "axios";
import noImg from '../../assets/no-img.png';
import { Buffer } from 'buffer';

export default class DatabaseFetcher {
    async getAllPhotos() {
        try {
            const response = await axios.get('http://backend:3000/files');
            const photos = response.data;
            
            // Create an array of promises for parallel execution
            const imgPromises = photos
                .map(async (photo: { url: string; }) => {
                    const photoId = photo.url.split('/').pop();
                    if (photoId && photoId.includes('example')) {
                        return noImg;
                    } else {
                        const imgResponse = await axios.get(
                            `http://backend:3000/files/${photoId}`,
                            { responseType: 'arraybuffer' }  // Get binary data
                        );
                        
                        // Convert binary data to base64
                        const base64 = Buffer.from(imgResponse.data).toString('base64');
                        const contentType = imgResponse.headers['content-type'];
                        return `data:${contentType};base64,${base64}`;
                    }
                });

            // Wait for all image requests to complete
            const imgs = await Promise.all(imgPromises);

            return {
                photos: photos,
                imgs: imgs  // Now contains array of base64 strings
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
            const response = await axios.post('http://backend:3000/upload', formData, {
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
}
