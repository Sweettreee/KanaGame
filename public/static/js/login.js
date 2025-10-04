let app = null;

export async function initLogin(container) {

    destroyLogin(); // 기존 앱이 있으면 제거

    // Pixi Application 생성
    app = new PIXI.Application();
    await app.init({
        width: 800,
        height: 600,
        backgroundColor: 0x1099bb
    });
    console.log(app);
    container.appendChild(app.canvas);
    
    // ------------------------
    // 배경
    // ------------------------
    const bg = new PIXI.Graphics();
    bg.beginFill(0xeeeeee);
    bg.drawRect(150, 100, 500, 400); // 로그인 박스
    bg.endFill();
    app.stage.addChild(bg);
    
    // ------------------------
    // 제목
    // ------------------------
    const title = new PIXI.Text("Login", {
        fontFamily: "Arial",
        fontSize: 36,
        fill: 0x333333
    });
    title.x = 400 - title.width / 2;
    title.y = 130;
    app.stage.addChild(title);
    
    // ------------------------
    // 아이디 / 비밀번호 텍스트
    // ------------------------
    const labelID = new PIXI.Text("ID:", { fontSize: 24, fill: 0x333333 });
    labelID.x = 200;
    labelID.y = 220;
    app.stage.addChild(labelID);
    
    const labelPW = new PIXI.Text("Password:", { fontSize: 24, fill: 0x333333 });
    labelPW.x = 200;
    labelPW.y = 280;
    app.stage.addChild(labelPW);
    
    // ------------------------
    // 입력 박스 (간단히 사각형으로)
    function createInputBox(x, y) {
        const box = new PIXI.Graphics();
        box.lineStyle(2, 0x333333);
        box.beginFill(0xffffff);
        box.drawRect(0, 0, 300, 30);
        box.endFill();
        box.x = x;
        box.y = y;
        return box;
    }
    
    const inputID = createInputBox(300, 220);
    const inputPW = createInputBox(300, 280);
    app.stage.addChild(inputID, inputPW);
    
    // ------------------------
    // 로그인 버튼
    // ------------------------
    const loginBtn = new PIXI.Graphics();
    loginBtn.beginFill(0x00cc66);
    loginBtn.drawRoundedRect(0, 0, 120, 40, 10);
    loginBtn.endFill();
    loginBtn.x = 340;
    loginBtn.y = 350;
    loginBtn.interactive = true;
    loginBtn.buttonMode = true;
    app.stage.addChild(loginBtn);
    
    const btnText = new PIXI.Text("Login", { fontSize: 24, fill: 0xffffff });
    btnText.x = loginBtn.x + (loginBtn.width - btnText.width) / 2;
    btnText.y = loginBtn.y + (loginBtn.height - btnText.height) / 2;
    app.stage.addChild(btnText);
    
    // ------------------------
    // 버튼 클릭 이벤트
    // ------------------------
    loginBtn.on("pointerdown", () => {
        alert("로그인 버튼 클릭됨!\n(실제 서버 연동은 구현 필요)");
    });
    
    // ------------------------
    // 간단히 입력값 구현
    // 실제 입력은 HTML input 태그를 겹쳐서 처리하는 방식이 일반적
    // Pixi로만 구현하면 키보드 이벤트 직접 처리 필요
    // ------------------------
}

export function destroyLogin() {
    if (app) {
        app.destroy(true, { children: true, texture: true, baseTexture: true });
        app = null;
    }
}