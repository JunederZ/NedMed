import {type Actions } from "@sveltejs/kit";
// import type {PageServerLoad} from "./$types.js";
import axios from "axios";

// export const load: PageServerLoad = (event) => {
//     event.cookies.set("sessionid", new Date().getTime().toString(), { path : ""})
//     return
// }

export const actions = {
    default: async (event) => {

        const data = await event.request.formData()
        const file = data.get("file") as File | null;

        if (!file?.size) {
            console.log("No file");
            return { success: false, message: "No file selected." };
        }

        const formData = new FormData();
        formData.append('image', file);

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
} satisfies Actions