<script lang="ts">
    import PhotoCard from "$lib/components/PhotoCard.svelte";
</script>

<div class="flex flex-col items-center h-screen bg-slate-950 text-white">
    <div class="mt-8">
        <h1 class="text-4xl font-bold text-center">Media Archive</h1>
        <p class="text-center">Keep your memories!</p>
    </div>
    <div class="flex h-[40vh] justify-center items-center gap-5">
       <PhotoCard />
       <PhotoCard />
    </div>
</div>