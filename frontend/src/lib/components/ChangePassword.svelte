<script lang="ts">
    import { toast } from 'svelte-sonner';

    let currentPassword = $state('');
    let newPassword = $state('');
    let confirmPassword = $state('');
    let isSubmitting = $state(false);
    let passwordError = $state('');

    // 密码验证规则
    const passwordRules = [
        { test: (p: string) => p.length >= 8, message: '密码长度至少需要8个字符' },
        { test: (p: string) => /[A-Z]/.test(p), message: '必须包含大写字母' },
        { test: (p: string) => /[a-z]/.test(p), message: '必须包含小写字母' },
        { test: (p: string) => /[0-9]/.test(p), message: '必须包含数字' },
        { test: (p: string) => /[!@#$%^&*(),.?":{}|<>]/.test(p), message: '必须包含特殊字符' }
    ];

    function validatePassword(password: string): string {
        for (const rule of passwordRules) {
            if (!rule.test(password)) {
                return rule.message;
            }
        }
        return '';
    }

    async function handleSubmit() {
        if (isSubmitting) return;

        // 重置错误
        passwordError = '';

        // 验证新密码
        passwordError = validatePassword(newPassword);
        if (passwordError) {
            toast.error(passwordError);
            return;
        }

        // 验证确认密码
        if (newPassword !== confirmPassword) {
            toast.error('两次输入的密码不一致');
            return;
        }

        isSubmitting = true;

        try {
            const response = await fetch('/api/auth/change-password', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify({
                    currentPassword,
                    newPassword
                })
            });

            if (!response.ok) {
                const error = await response.json();
                throw new Error(error.message || '修改密码失败');
            }

            toast.success('密码修改成功');
            currentPassword = '';
            newPassword = '';
            confirmPassword = '';
        } catch (error: unknown) {
            toast.error(error instanceof Error ? error.message : '修改密码失败');
        } finally {
            isSubmitting = false;
        }
    }
</script>

<div class="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-6 text-gray-800">修改密码</h2>
    
    <form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="space-y-4">
        <div>
            <label for="currentPassword" class="block text-sm font-medium text-gray-700">
                当前密码
            </label>
            <input
                type="password"
                id="currentPassword"
                bind:value={currentPassword}
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                required
            />
        </div>

        <div>
            <label for="newPassword" class="block text-sm font-medium text-gray-700">
                新密码
            </label>
            <input
                type="password"
                id="newPassword"
                bind:value={newPassword}
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                required
            />
        </div>

        <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
                确认新密码
            </label>
            <input
                type="password"
                id="confirmPassword"
                bind:value={confirmPassword}
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
                required
            />
        </div>

        <!-- 密码规则提示 -->
        <div class="text-sm text-gray-600 space-y-1">
            <p>密码必须满足：</p>
            <ul class="list-disc pl-5">
                <li>至少8个字符</li>
                <li>包含大写字母</li>
                <li>包含小写字母</li>
                <li>包含数字</li>
                <li>包含特殊字符</li>
            </ul>
        </div>

        <button
            type="submit"
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            disabled={isSubmitting}
        >
            {isSubmitting ? '提交中...' : '修改密码'}
        </button>
    </form>
</div> 