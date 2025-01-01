import numeral from 'numeral'
import dayjs from 'dayjs'

export const formatNumber = (num: number, format: string = '0,0') => {
	return numeral(num).format(format)
}

export const ignoreFindDOMNodeError = () => {
	// eslint-disable-next-line
	const consoleError = console.error.bind(console)
	// eslint-disable-next-line
	console.error = (errObj, ...args) => {
		if (typeof errObj === 'string' && args.includes('findDOMNode')) {
			return
		}
		consoleError(errObj, ...args)
	}
}

export const dateTimeFormatter = (date, format = "MM-DD-YYYY") => dayjs(date).format(format)

export const INPUT_TRIM = {
	getValueFromEvent: e => e?.target.value?.trim()
}

export const INPUT_LETTERS = {
	getValueFromEvent: e => {
		const trimmedValue = e?.target.value?.trim()
		const nonNumericValue = trimmedValue.replace(/[0-9]/g, '')
		return nonNumericValue
	}
}

export const INPUT_EMAIL = {
	getValueFromEvent: e => {
		const trimmedValue = e?.target.value?.trim();
		const validEmail = trimmedValue.replace(/[^\w.@+-]/g, ''); // Remove characters not allowed in an email address
		return validEmail;
	}
}

export const INPUT_NUMBER = {
	getValueFromEvent: e => {
		const trimmedValue = e?.target.value?.trim()
		const numericValue = trimmedValue.replace(/\D/g, '') // Allow only numeric characters
		return numericValue
	}
}
