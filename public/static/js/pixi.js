import { SpineBoyCharacter } from "./PixiGame/spineBoy.js";
let app = null;

export async function pixiGame (container) {
    (async () => {
        // Create a PixiJS application.
        app = new PIXI.Application();
        await app.init({
            background: '#1099bb', resizeTo: window
        });
        console.log(app);
        container.appendChild(app.canvas);
        // Load the assets.
        await PIXI.Assets.load([
            {
            alias: 'spineSkeleton',
            src: 'https://raw.githubusercontent.com/pixijs/spine-v8/main/examples/assets/spineboy-pro.skel',
            },
            {
            alias: 'spineAtlas',
            src: 'https://raw.githubusercontent.com/pixijs/spine-v8/main/examples/assets/spineboy-pma.atlas',
            },
            {
            alias: 'sky',
            src: 'https://pixijs.com/assets/tutorials/spineboy-adventure/sky.png',
            },
            {
            alias: 'background',
            src: 'https://pixijs.com/assets/tutorials/spineboy-adventure/background.png',
            },
            {
            alias: 'midground',
            src: 'https://pixijs.com/assets/tutorials/spineboy-adventure/midground.png',
            },
            {
            alias: 'platform',
            src: 'https://pixijs.com/assets/tutorials/spineboy-adventure/platform.png',
            },
        ]);
        const spineBoy = await SpineBoyCharacter(app);
        app.stage.addChild(spineBoy);
        
    })();
}