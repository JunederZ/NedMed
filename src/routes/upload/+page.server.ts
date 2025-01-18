import type {Actions} from "@sveltejs/kit";
import type {PageServerLoad} from "./$types.js";

export const load: PageServerLoad = (event) => {
    console.log("Load page server...");
    event.cookies.set("sessionid", new Date().getTime().toString(), { path : ""})
    return
}

export const actions = {
    upload: async (event) => {

        const data = event.request.formData()
        const file = data.then(value => {
            return value.get('file')
            })

        const formData = new FormData();
        formData.append("file", file);

    }
} satisfies Actions