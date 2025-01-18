import type {Actions} from "@sveltejs/kit";
import type {PageServerLoad} from "./$types.js";

export const load: PageServerLoad = (event) => {
    console.log("Load page server...");
    event.cookies.set("sessionid", new Date().getTime().toString(), { path : ""})
    return
}

export const actions = {
    upload: async (event) => {

        const data = await event.request.formData()
        const file = data.get("file") as File | null;

        if (!file) {
            console.log("No file");
            return { success: false, message: "No file selected." };
        }


        const formData = new FormData();
        formData.append('image', file);
        console.log(formData);


        try {
            const response = await fetch('http://127.0.0.1:3000/upload', {
                method: 'POST',
                body: formData,
            });

            if (response.ok) {
                console.log("Successfully uploaded!");
                return { success: true };
            } else {
                console.log("Failed to upload" + await response.text());
                return { success: false, message: 'Upload failed.' + (await response.text()) };
            }
        } catch (error) {
            console.error("Error during upload:", error);
            return { success: false, message: 'Upload failed.' };
        }



    }
} satisfies Actions