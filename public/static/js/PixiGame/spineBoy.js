let app = null;

export async function SpineBoyCharacter (app) {
    const spineBoy = spine.Spine.from({
        skeleton: 'spineSkeleton',
        atlas: 'spineAtlas',
        scale: 0.5,
        autoUpdate: true,
    });

    spineBoy.x = app.screen.width / 2;
    spineBoy.y = app.screen.height - 80;
    spineBoy.scale.set(0.5);
    
    return spineBoy;

}