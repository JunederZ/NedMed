<script lang="ts">
    import PhotoCard from "$lib/components/PhotoCard.svelte";
    import type { PageData } from "./$types.js";

    let { data }: { data: PageData } = $props()

    console.log(data)

</script>

<div class="flex flex-col items-center h-screen w-screen bg-slate-950 text-white">
    <div class="mt-8">
        <h1 class="text-4xl font-bold text-center">Media Archive</h1>
        <p class="text-center">Keep your memories!</p>
    </div>
    <div class="grid grid-cols-[repeat(auto-fit,_16.666666%)] place-items-center m-10 gap-10 w-screen justify-center items-center">
        {#each data.photos as photo}
            <PhotoCard {...photo}/>
        {/each}
    </div>
</div>