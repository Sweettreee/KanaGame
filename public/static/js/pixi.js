import * as PIXI from "https://pixijs.download/release/pixi.mjs";

const app = new PIXI.Application({ width: 800, height: 600, backgroundColor: 0x1099bb });
document.getElementById('game').appendChild(app.view);

const bunny = PIXI.Sprite.from('https://pixijs.com/assets/bunny.png');
bunny.anchor.set(0.5);
bunny.x = app.renderer.width / 2;
bunny.y = app.renderer.height / 2;
app.stage.addChild(bunny);

app.ticker.add(() => {
    bunny.rotation += 0.01;
});