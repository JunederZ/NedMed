import type {PageServerLoad} from "../../.svelte-kit/types/src/routes/$types.js";
import DatabaseFetcher from "$lib/server/database/fetch.js";

export const load: PageServerLoad = async () => {
    try {
        const fetch = new DatabaseFetcher();
        const { photos, imgs } = await fetch.getAllPhotos();
        
        return {
            photoData: photos,
            imgData: imgs
        };
    } catch (error) {
        console.error('Error in load function:', error);

        return {
            photoData: [],
            imgData: []
        };
    }
};
