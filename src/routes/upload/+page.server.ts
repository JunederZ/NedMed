import {type Actions } from "@sveltejs/kit";
// import type {PageServerLoad} from "./$types.js";
// import axios from "axios";
import DatabaseFetcher from "../../lib/server/database/fetch.js"


// export const load: PageServerLoad = (event) => {
//     event.cookies.set("sessionid", new Date().getTime().toString(), { path : ""})
//     return
// }

export const actions = {
    default: async (event) => {

        const data = await event.request.formData()
        const file = data.get("file") as File | null;
        const title = data.get("title") as string;
        const description = data.get("description") as string;

        if (!file?.size) {
            console.log("No file");
            return { success: false, message: "No file selected." };
        }

        const fetch = new DatabaseFetcher()
        return await fetch.uploadPhoto(file, title, description)

    }
} satisfies Actions