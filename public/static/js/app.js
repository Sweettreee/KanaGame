export async function CreateApp() {
    (async () => {
        const Down = [
            await PIXI.Assets.load('../../images/MoveDown1.png'),
            await PIXI.Assets.load('../../images/MoveDown2.png'),
            await PIXI.Assets.load('../../images/MoveDown3.png')
        ];

        const Up = [
            await PIXI.Assets.load('../../images/MoveUp1.png'),
            await PIXI.Assets.load('../../images/MoveUp2.png'),
            await PIXI.Assets.load('../../images/MoveUp3.png')
        ];

        const Left = [
            await PIXI.Assets.load('../../images/MoveLeft1.png'),
            await PIXI.Assets.load('../../images/MoveLeft2.png'),
            await PIXI.Assets.load('../../images/MoveLeft3.png')
        ]

        const direction = { x: 0, y: 0 };
        let changedTextre = Down[0];
        // Set up keyboard event listeners
        window.addEventListener('keydown', (e) => {
            switch(e.key) {
                case 'w':
                    direction.y = -1;
                    changedTextre = Up[0];
                    break;
                case 's':
                    direction.y = 1;
                    changedTextre = Down[0];
                    break;
                case 'a':
                    direction.x = -1;
                    changedTextre = Left[0];
                    break;
                case 'd':
                    direction.x = 1;
                    break;
            }
        });

        window.addEventListener('keyup', (e) => {
            switch(e.key) {
                case 'w':
                    changedTextre = Down[0];
                    direction.y = 0;
                    direction.x = 0;
                    break;
                case 's':
                    changedTextre = Down[0];
                    direction.y = 0;
                    direction.x = 0;
                    break;
                case 'a':
                    changedTextre = Down[0];
                    direction.y = 0;
                    direction.x = 0;
                    break;
                case 'd':
                    changedTextre = Down[0];
                    direction.y = 0;
                    direction.x = 0;
                    break;
            }
        });
        const app = new PIXI.Application();

        await app.init({ background: '#1099bb', resizeTo: window });

        document.body.appendChild(app.canvas);

        const container = new PIXI.Container();

        app.stage.addChild(container);

        const texture = Down[0];
        const pokemon = new PIXI.Sprite(texture);

        // 캐릭터 설정
        pokemon.width = 50;
        pokemon.height = 50;

        pokemon.anchor.set(0.5);
        pokemon.x = app.screen.width / 2;
        pokemon.y = app.screen.height / 2;

        container.addChild(pokemon);

        const speed = 5;

        app.ticker.add(() => {
            pokemon.x += direction.x * speed;
            pokemon.y += direction.y * speed;
            pokemon.texture = changedTextre;
        });
    })();
}

