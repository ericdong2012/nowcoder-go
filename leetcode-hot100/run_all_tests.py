"""跑通 leetcode-hot100 全部自测：python run_all_tests.py"""

import os
import runpy
import sys

DIR = os.path.dirname(os.path.abspath(__file__))


def main():
    passed, failed = [], []
    for name in sorted(os.listdir(DIR)):
        if not name.endswith(".py") or not name[0].isdigit():
            continue
        try:
            runpy.run_path(os.path.join(DIR, name), run_name="__main__")
            passed.append(name)
        except Exception as e:  # noqa: BLE001
            failed.append((name, repr(e)))

    print(f"\n===== 通过 {len(passed)}/100 =====")
    if failed:
        print("失败清单：")
        for name, err in failed:
            print(f"  {name}: {err[:120]}")
    return 1 if failed else 0


if __name__ == "__main__":
    # 加 main 保护：被其它工具 import 时不会触发全量测试并 sys.exit 劫持宿主进程
    sys.exit(main())
