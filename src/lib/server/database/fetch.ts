
import axios from "axios";

export default class DatabaseFetcher {
  async getAllPhotos() {
    return axios.get('http://localhost:3000/files')
    .then(response => {
        return response.data;
    });
  }
  
  // async getPhoto(id: string) {
  //   return axios.get(`http://localhost:3000/files/${id}`)
  //   .then(response => {
  //       return response.data;
  //   });
  // }

  async uploadPhoto(photo: Blob, title: string, description: string) {
    const formData = new FormData();
    formData.append('image', photo);
    formData.append('title', title);
    formData.append('description', description);

    try {
        const response = await axios.post('http://127.0.0.1:3000/upload', formData, {
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